package twitter

import (
	"GIG/app/storages"
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

type UploadSuccessResponse struct {
	MediaId int `json:"media_id" bson:"media_id"`
}

func UploadMedia(imageUrlString string) (mediaId int, err error) {
	imageUrl := strings.Split(imageUrlString, "/")
	if len(imageUrl) == 3 { // if a valid image exist
		title := imageUrl[1]
		filename := imageUrl[2]

		localFile, err := storages.FileStorageHandler{}.GetFile(title, filename)
		if err != nil {
			return 0, err
		}

		method := "POST"
		payload := &bytes.Buffer{}
		writer := multipart.NewWriter(payload)
		part1, errFile1 := writer.CreateFormFile("media", filename)
		if errFile1 != nil {
			return 0, errFile1
		}
		_, errFile1 = io.Copy(part1, localFile)
		if errFile1 != nil {
			return 0, errFile1
		}
		err = writer.Close()
		if err != nil {
			return 0, err
		}

		req, err := http.NewRequest(method, MediaUploadUrl, payload)

		if err != nil {
			return 0, err
		}

		client := &http.Client{
		}

		req.Header.Add("Authorization", GetAuthHeader())
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
			json.Unmarshal(body,&errorBody)
			log.Println(errorBody)
			return 0, errors.New("error uploading media to twitter.")
		}

		var jsonBody UploadSuccessResponse
		json.Unmarshal(body, &jsonBody)
		return jsonBody.MediaId, nil
	}
	return 0, errors.New("invalid image url.")
}
