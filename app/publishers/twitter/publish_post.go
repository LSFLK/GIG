package twitter

import (
	"GIG-SDK/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func PublishPost(entity models.Entity, mediaId int) error {
	url := PublishPostUrl + "?status=" + entity.Snippet
	if mediaId != 0 {
		url = url + "&media_ids=" + string(mediaId)
	}
	method := "POST"

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}
	req.Header.Add("Authorization", GetAuthHeader())

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
		log.Println(errorBody)
		return errors.New("error posting to twitter")
	}

	// TODO: not publishing post, search entity names on twitter (twitter handles)


	return nil
}
