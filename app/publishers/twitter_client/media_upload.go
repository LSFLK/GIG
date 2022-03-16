package twitter_client

import (
	"GIG/app/publishers/twitter_client/functions"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type UploadSuccessResponse struct {
	MediaId int `json:"media_id" bson:"media_id"`
}

func UploadMedia(imageUrlString string) (mediaId int, err error) {
	err, title, filename := functions.GetTitleAndFilenameFromUrl(imageUrlString)

	if err != nil {
		return 0, err
	}

	err, payload, writer := functions.CreatePayload(title, filename)
	if err != nil {
		return 0, err
	}

	method := "POST"
	client := GetHttpClient()
	req, err := http.NewRequest(method, MediaUploadUrl, payload)

	if err != nil {
		return 0, err
	}

	//req.Header.Add("Authorization", GetAuthHeader())
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	if res.StatusCode != 200 {
		var errorBody interface{}
		json.Unmarshal(body, &errorBody)
		return 0, errors.New("error uploading media to twitter. " + fmt.Sprintf("%v", errorBody))
	}

	var jsonBody UploadSuccessResponse
	json.Unmarshal(body, &jsonBody)
	return jsonBody.MediaId, nil
}
