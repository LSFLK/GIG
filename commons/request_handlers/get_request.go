package request_handlers

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
)

const requestHeaderKey = "User-Agent"
const requestHeaderValue = "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"

/**
	get the response string for a given url
 */
func GetRequest(uri string) (string, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := http.Client{Transport: transport, Timeout: 30 * time.Second}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set(requestHeaderKey, requestHeaderValue)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, bodyError := ioutil.ReadAll(resp.Body)
	if bodyError != nil {
		return "", bodyError
	}

	return string(body), nil
}
