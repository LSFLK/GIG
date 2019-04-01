// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/models"
	"GIG/app/utility/decoders"
	"GIG/app/utility/requesthandlers"
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/collectlinks"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	bson2 "gopkg.in/mgo.v2/bson"
	"io"
	"log"
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
	decoder := decoders.WikipediaDecoder{}
	queue := make(chan string)
	go func() { queue <- args[0] }()

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

	for uri := range queue {
		response := enqueue(uri, queue)
		entity := decoder.DecodeSource(response)
		entity.ID=bson2.ObjectId(uri)

		var result models.Entity
		db.Collection("entities").FindOne(context.TODO(), bson.M{"_id": uri}).Decode(&result)
		if string(result.ID) == "" {
			insertResult, err := db.Collection("entities").InsertOne(context.TODO(), entity)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		}

	}
}

func enqueue(uri string, queue chan string) *bytes.Buffer {
	fmt.Println("fetching", uri)
	visited[uri] = true

	client, req := requesthandlers.SendRequest("GET", uri)

	resp, err := client.Do(req)

	if err != nil {
		return &bytes.Buffer{}
	}
	var bufferedResponse bytes.Buffer
	response := io.TeeReader(resp.Body, &bufferedResponse)
	links := collectlinks.All(response)
	defer resp.Body.Close()

	for _, link := range links {
		absolute := fixUrl(link, uri)
		if uri != "" {
			if !visited[absolute] {
				go func() { queue <- absolute }()
			}
		}
	}
	return &bufferedResponse
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
