package twitter_client

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"net/http"
)

func GetTwitterClient() *twitter.Client{
	// Twitter client
	return twitter.NewClient(GetHttpClient())
}

func GetHttpClient() *http.Client{
	config := oauth1.NewConfig(ConsumerKey, ConsumerSecret)
	token := oauth1.NewToken(AccessToken, TokenSecret)
	// http.Client will automatically authorize Requests
	return config.Client(oauth1.NoContext, token)
}
