package publishers

import "GIG/app/publishers/twitter"

func LoadPublishers() {
	twitter.LoadTwitter()
}
