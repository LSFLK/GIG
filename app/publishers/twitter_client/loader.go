package twitter_client

import "github.com/revel/revel"

var (
	SignatureMethod string
	ConsumerKey     string
	ConsumerSecret  string
	AccessToken     string
	TokenSecret     string
	AuthSignature   string
)

func LoadTwitter() {
	SignatureMethod, _ = revel.Config.String("twitter.signatureMethod")
	ConsumerKey, _ = revel.Config.String("twitter.consumerKey")
	ConsumerSecret, _ = revel.Config.String("twitter.consumerSecret")
	AccessToken, _ = revel.Config.String("twitter.accessToken")
	TokenSecret, _ = revel.Config.String("twitter.tokenSecret")
	AuthSignature, _ = revel.Config.String("twitter.authSignature")
}