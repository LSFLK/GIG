package request_handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/**
Post to an url with data
 */
func PostRequest(uri string, data interface{}) (string, error) {

	// json encode interface
	b, err := json.Marshal(data)
	var jsonStr = []byte(b)

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, bodyError := ioutil.ReadAll(resp.Body)
	if bodyError != nil {
		return "", bodyError
	}

	return string(body), err
}
