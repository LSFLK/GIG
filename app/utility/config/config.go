package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var configs map[string]string

func GetConfig(configName string) string {
	if len(configs)==0 {
		fmt.Println("Loading utility configuration...")
		file, err := os.Open("conf/config.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&configs)
		if err != nil {
			fmt.Println("Error loading utility configuration!")
			panic(err)
		}
	}

	return configs[configName]
}
