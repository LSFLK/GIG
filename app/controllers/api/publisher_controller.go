package api

import (
	"GIG-SDK/models"
	"GIG/app/controllers"
	"GIG/app/storages"
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/revel/revel"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type PublisherController struct {
	*revel.Controller
}

const (
	chunksize int = 1024
)

func FileToBuffer(data *os.File) (buffer *bytes.Buffer) {

	var (
		part  []byte
		err   error
		count int
	)

	reader := bufio.NewReader(data)
	buffer = bytes.NewBuffer(make([]byte, 0))
	part = make([]byte, chunksize)

	for {
		if count, err = reader.Read(part); err != nil {
			break
		}
		buffer.Write(part[:count])
	}
	if err != io.EOF {
		log.Println("Error Reading ", data.Name(), ": ", err)
	}

	return
}

func (c PublisherController) Twitter() revel.Result {
	var (
		err    error
		entity models.Entity
	)
	log.Println("twitter publish request")
	err = c.Params.BindJSON(&entity)

	if err != nil {
		log.Println("binding error:", err)
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrorResponse(err, 403))
	}

	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrorResponse(err, 500))
	}

	imageUrl := strings.Split(entity.ImageURL, "/")
	if len(imageUrl) == 3 { // if a valid image exist
		title := imageUrl[1]
		filename := imageUrl[2]

		localFile, err := storages.FileStorageHandler{}.GetFile(title, filename)
		if err != nil {
			log.Println(err)
			c.Response.Status = 400
			return c.RenderJSON(err)
		}

		uploadUrl := "https://upload.twitter.com/1.1/media/upload.json"
		method := "POST"
		payload := &bytes.Buffer{}
		writer := multipart.NewWriter(payload)
		part1, errFile1 := writer.CreateFormFile("media", filename)
		if errFile1 != nil {
			log.Println(errFile1)
		}
		_, errFile1 = io.Copy(part1, localFile)
		if errFile1 != nil {
			log.Println(errFile1)
		}
		err = writer.Close()
		if err != nil {
			log.Println(err)
		}

		req, err := http.NewRequest(method, uploadUrl, payload)

		if err != nil {
			log.Println(err)
			return c.RenderJSON(controllers.BuildErrorResponse(err, 500))
		}

		client := &http.Client{
		}

		req.Header.Add("Authorization", "OAuth oauth_consumer_key=\"EocLQDL2P11GgTGSfbezhC41V\",oauth_token=\"1415595042421411845-pFZZ2WI9Hu70KNcMaqNGhyoOkkEhFn\",oauth_signature_method=\"HMAC-SHA1\",oauth_timestamp=\"1645602224\",oauth_nonce=\"p1s1c2yD0JI\",oauth_version=\"1.0\",oauth_signature=\"zP98hwUovG%2B5P6rvtzsA8rp%2FexE%3D\"")
		req.Header.Add("Cookie", "personalization_id=\"v1_Jn7MfA+YJSgdnWfVjrj0Dg==\"; guest_id=v1%3A162695026803727867; guest_id_marketing=v1%3A162695026803727867; guest_id_ads=v1%3A162695026803727867")

		req.Header.Set("Content-Type", writer.FormDataContentType())
		res, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return c.RenderJSON(controllers.BuildErrorResponse(err, 500))
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Println("error uploading media to twitter.", res.StatusCode)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return c.RenderJSON(controllers.BuildErrorResponse(err, 500))
		}

		var jsonBody interface{}
		json.Unmarshal(body, &jsonBody)
		return c.RenderJSON(controllers.BuildSuccessResponse(jsonBody, 200))
	}

	return c.RenderJSON(controllers.BuildSuccessResponse("publish request queued.", 200))

}
