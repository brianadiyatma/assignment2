package model

type Response struct {
	ResponseCode int         `json:"response_code"`
	Success      bool        `json:"success"`
	Error        interface{} `json:"error"`
	Data         interface{} `json:"data"`
}