package config

import (
	"encoding/json"
	"os"
)

func GetConfig(configName string) string {
	file, err := os.Open("conf/config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var configs map[string]string
	err = decoder.Decode(&configs)
	if err != nil {
		panic(err)
	}
	return configs[configName]
}
