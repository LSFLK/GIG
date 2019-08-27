package utils

import (
	"GIG/commons/request_handlers"
	"GIG/scripts"
	"encoding/json"
)

type NERResult struct {
	Category   string
	EntityName string
}

/**
NER extraction
 */
func ExtractEntityNames(textContent string) ([]NERResult, error) {

	apiResp, err := request_handlers.PostRequest(scripts.NERServerUrl, textContent)
	if err != nil {
		return nil, err
	}

	var (
		entityTitles [][]string
		results      []NERResult
	)
	if err = json.Unmarshal([]byte(apiResp), &entityTitles);err!=nil{
		return nil, err
	}

	for _, entity := range entityTitles {
		newNERResult := NERResult{EntityName: entity[0], Category: entity[1]}
		results = append(results, newNERResult)
	}
	return results, nil
}
