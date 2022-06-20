package controllers

import (
	"github.com/lsflk/gig-sdk/client"
)

/*
ServerUrl - Set the GIG server API url here for crawlers
*/
const (
	ServerUrl = "http://localhost:9000/"
)

var testClient = client.GigClient{
	ApiUrl:                 ServerUrl + "api/",
	ApiKey:                 "[ApiKey]",
	NerServerUrl:           "http://localhost:8081/classify",
	NormalizationServerUrl: ServerUrl + "api/",
	OcrServerUrl:           "http://localhost:8082/extract?url=",
}
