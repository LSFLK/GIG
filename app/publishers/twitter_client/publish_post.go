package twitter_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lsflk/gig-sdk/models"
	"io/ioutil"
	"net/http"
)

func PublishPost(entity models.Entity, mediaId int) error {
	tweetUrl := CreateTweet(entity, mediaId)
	method := "POST"

	client := GetHttpClient()
	req, err := http.NewRequest(method, tweetUrl, nil)

	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		var errorBody interface{}
		json.Unmarshal(body, &errorBody)
		return errors.New("error publishing post to twitter. " + fmt.Sprintf("%v", errorBody))
	}

	return nil
}
