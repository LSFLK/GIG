package controllers

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type Controller map[string]interface{}

func ParseUintOrDefault(intStr string, _default uint64) uint64 {
	if value, err := strconv.ParseUint(intStr, 0, 64); err != nil {
		return _default
	} else {
		return value
	}
}

func ParseIntOrDefault(intStr string, _default int64) int64 {
	if value, err := strconv.ParseInt(intStr, 0, 64); err != nil {
		return _default
	} else {
		return value
	}
}

func ConvertToObjectIdHex(id string) (result bson.ObjectId, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("unable to convert %v to object id", id)
		}
	}()

	return bson.ObjectIdHex(id), err
}

func BuildErrResponse( status int, err error) Controller {
	ctrlErr := Controller{}
	ctrlErr["error_message"] = err.Error()
	ctrlErr["status"] = status
	return ctrlErr
}

func BuildResponse(status int, content interface{}) Controller {
	ctrlResp := Controller{}
	ctrlResp["status"] = status
	ctrlResp["content"] = content
	return ctrlResp
}
