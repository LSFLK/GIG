package twitter

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

func GetAuthHeader() string {
	return "OAuth oauth_consumer_key=\"" + ConsumerKey +
		"\",oauth_token=\"" + AccessToken +
		"\",oauth_signature_method=\"" + SignatureMethod +
		"\",oauth_timestamp=\"1645619155\",oauth_nonce=\"66UGLwrlFro\",oauth_version=\"1.0\",oauth_signature=\"" + AuthSignature + "\""
}
