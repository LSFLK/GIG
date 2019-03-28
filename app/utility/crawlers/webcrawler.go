// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/utility/requesthandlers"
	"bytes"
	"flag"
	"fmt"
	"github.com/collectlinks"
	"io"
	"io/ioutil"
	"net/url"
	"os"
)

var visited = make(map[string]bool)

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("starting url not specified")
		os.Exit(1)
	}
	queue := make(chan string)
	go func() { queue <- args[0] }()
	for uri := range queue {
		body := enqueue(uri, queue)
		fmt.Println(string(body))
	}
}

func enqueue(uri string, queue chan string) []byte {
	fmt.Println("fetching", uri)
	visited[uri] = true

	client, req := requesthandlers.SendRequest("GET", uri)

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	var bufferedResponse bytes.Buffer
	response := io.TeeReader(resp.Body, &bufferedResponse)

	body, err := ioutil.ReadAll(response)
	links := collectlinks.All(&bufferedResponse)

	defer resp.Body.Close()

	for _, link := range links {
		absolute := fixUrl(link, uri)
		if uri != "" {
			if !visited[absolute] {
				go func() { queue <- absolute }()
			}
		}
	}

	return body

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
