package requesthandlers

import (
	"crypto/tls"
	"net/http"
	"time"
)

const requestHeaderKey = "User-Agent"
const requestHeaderValue = "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"

func GetRequest(uri string) (*http.Response, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := http.Client{Transport: transport, Timeout: 10 * time.Second}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set(requestHeaderKey, requestHeaderValue)
	resp, err := client.Do(req)

	return resp, err
}
