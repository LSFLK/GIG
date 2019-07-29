package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ServerApiUrl string
	MapApiUrl    string
	AppKey       string
}

func GetConfig() Config {
	file, err := os.Open("conf/config.json")
	if err != nil {
		fmt.Println("file error:", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("utility config declaration error:", err)
	}
	return configuration
}
