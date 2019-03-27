// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/models"
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/collectlinks"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	bson2 "gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
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
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := http.Client{Transport: transport}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

	//TODO : sending get request two times because response is modified when using ioutil.ReadAll reduce it to one
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	resp2, err := client.Do(req)

	body, err := ioutil.ReadAll(resp2.Body)
	links := collectlinks.All(resp.Body)

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	mongoClient.Connect(nil)
	err = mongoClient.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	db := mongoClient.Database("gig")
	entity := models.Entity{
		ID:        bson2.NewObjectId(),
		Title:     uri,
		Content:   string(body),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	var result models.Entity
	db.Collection("entities").FindOne(context.TODO(), bson.M{"_id": uri}).Decode(&result)
	if string(result.ID) == "" {
		insertResult, err := db.Collection("entities").InsertOne(context.TODO(), entity)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	}

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
