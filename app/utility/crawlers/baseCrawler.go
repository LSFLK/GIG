// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/utility/requestHandlers"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/collectlinks"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var visited = make(map[string]bool)

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}
	queue := make(chan string)
	go func() { queue <- args[0] }()
	for uri := range queue {
		enqueue(uri, queue)
	}
}


func enqueue(uri string, queue chan string) {
	fmt.Println("fetching", uri)
	visited[uri] = true

	client, req:= requestHandlers.SendRequest("GET", uri)

	/**
	TODO : fix error: sending get request two times because response
	is modified when using ioutil.ReadAll reduce it to one
	 */

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	resp2, err := client.Do(req)

	body, err := ioutil.ReadAll(resp2.Body)
	links := collectlinks.All(resp.Body)

	defer resp.Body.Close()
	defer resp2.Body.Close()

	for _, link := range links {
		absolute := fixUrl(link, uri)
		if uri != "" {
			if !visited[absolute] {
				go func() { queue <- absolute }()
			}
		}
	}

}

func fixUrl(href, base string) (string) {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}
