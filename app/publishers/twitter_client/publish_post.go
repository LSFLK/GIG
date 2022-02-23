package twitter_client

import (
	"GIG-SDK/models"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"strconv"
)

func PublishPost(entity models.Entity, mediaId int) error {
	url := PublishPostUrl + "?status=" + url2.QueryEscape(entity.Title+" #kavudalk view more at https://kavuda.lk/#/profile/"+url2.QueryEscape(entity.Title))
	if mediaId != 0 {
		url = url + "&media_ids=" + strconv.Itoa(mediaId)
	}
	method := "POST"
	log.Println(mediaId, entity.Title)
	client := GetHttpClient()
	req, err := http.NewRequest(method, url, nil)

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
		log.Println(string(body))
		return errors.New("error posting to twitter")
	}

	// TODO: search entity names on twitter (twitter handles)

	return nil
}
