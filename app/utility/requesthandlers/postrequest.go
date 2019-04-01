package requesthandlers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func PostRequest(uri string, data interface{}) (*http.Response, error) {

	// json encode interface
	b, err := json.Marshal(data)
	var jsonStr = []byte(b)

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	return resp, err
}
