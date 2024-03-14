package util

import "second-assignment/model"

func CreateResponse(responseCode int, success bool, err interface{}, data interface{}) model.Response {
	res := model.Response{ResponseCode: responseCode,Success:success, Error: err, Data: data }
	return res
}