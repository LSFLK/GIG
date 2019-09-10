package utils

import (
    "GIG/commons/request_handlers"
    "GIG/scripts"
    "encoding/json"
    "fmt"
    "github.com/revel/revel"
    "net/url"
    "strconv"
)

type Response struct {
    Status int `json:"status"`
    Content string `json:"content"`
}

/**
normalize entity title before appending
 */
func NormalizeName(title string) (string, error) {

    normalizedName, err := request_handlers.GetRequest(scripts.NormalizeServer + "?searchText=" + url.QueryEscape(title))

    if err != nil {
        return "", err
    }
    var response Response
    if json.Unmarshal([]byte(normalizedName), &response); err != nil {
        return "", err
    }
    fmt.Println(response)
    if response.Status != 200 {
        return "", revel.NewErrorFromPanic("Server responded with" + strconv.Itoa(response.Status))
    }
    fmt.Println("normalized", title, "to", response.Content)
    return response.Content, nil
}
