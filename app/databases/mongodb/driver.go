package mongodb

import (
	"log"
	"os"
)

var MaxPool int
var PATH string
var DBNAME string

func CheckAndInitServiceConnection() {
	if service.baseSession == nil {
		service.URL = PATH
		err := service.New()
		if err != nil {
			log.Println("error connecting to MongoDB database server:", service.URL)
			os.Exit(1)
		}
	}
}
