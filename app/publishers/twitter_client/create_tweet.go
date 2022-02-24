package twitter_client

import (
	"GIG-SDK/models"
	"net/url"
	"strconv"
)

func CreateTweet(entity models.Entity, mediaId int) (tweetUrl string) {
	tweetUrl = PublishPostUrl + "?status=" + url.QueryEscape(entity.Title+" - view more at "+KavudaLkUrl+url.QueryEscape(entity.Title))
	if mediaId != 0 {
		tweetUrl = tweetUrl + "&media_ids=" + strconv.Itoa(mediaId)
	}
	return
}
