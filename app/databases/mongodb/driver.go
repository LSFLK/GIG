package mongodb

import (
	"log"
	"os"
)

func CheckAndInitServiceConnection() {
	if service.baseSession == nil {
		service.URL = PATH
		err := service.New()
		if err != nil {
			log.Println("error connecting to MongoDB database server:", service.URL)
			log.Println(err)
			os.Exit(1)
		}
	}
}
