package main

import "github.com/rekkid/monitor/util/zjlog"

type ApiBody struct {
	Url     string `json:"url"`
	Method  string `json:"method"`
	ReqBody string `json:"req_body"`
}

func NewApiBody(url string, method string, reqBody string) *ApiBody {
	return &ApiBody{Url: url, Method: method, ReqBody: reqBody}
}

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

var (
	ErrRequestNotRe           = Err{Error: "api not recognized, bad request", ErrorCode: "001"}
	ErrRequestBodyParseFailed = Err{Error: "request body is not correct", ErrorCode: "002"}
	ErrInternalFults          = Err{Error: "internal service error", ErrorCode: "003"}
)

var log *zjlog.Log
