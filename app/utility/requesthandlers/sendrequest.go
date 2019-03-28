package requesthandlers

import (
	"crypto/tls"
	"net/http"
)

const requestHeaderKey = "User-Agent"
const requestHeaderValue = "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"

func SendRequest(method string, uri string) (http.Client, *http.Request) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := http.Client{Transport: transport}
	req, _ := http.NewRequest(method, uri, nil)
	req.Header.Set(requestHeaderKey, requestHeaderValue)

	return client, req
}
