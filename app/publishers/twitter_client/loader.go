package twitter_client

import "github.com/revel/revel"

var (
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	TokenSecret    string
	KavudaLkUrl    string
)

func LoadTwitter() {
	ConsumerKey, _ = revel.Config.String("twitter.consumerKey")
	ConsumerSecret, _ = revel.Config.String("twitter.consumerSecret")
	AccessToken, _ = revel.Config.String("twitter.accessToken")
	TokenSecret, _ = revel.Config.String("twitter.tokenSecret")
	KavudaLkUrl, _ = revel.Config.String("kavudaLk.webUrl")
}
