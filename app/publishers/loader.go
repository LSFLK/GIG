package publishers

import (
	"GIG/app/publishers/twitter_client"
)

func LoadPublishers() {
	twitter_client.LoadTwitter()
}
