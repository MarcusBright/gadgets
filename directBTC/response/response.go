package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	ErrorCode
	Data interface{} `json:"data,omitempty"`
}

// Code = 0 ok
// Code = -1 failed
// Code = -2 params error
type ErrorCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *ErrorCode) Error() string {
	return e.Msg
}

func ParamError(msg string) *ErrorCode {
	return &ErrorCode{
		Code: -2,
		Msg:  msg,
	}
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		switch e := err.(type) {
		case *ErrorCode:
			body.Code = e.Code
			body.Msg = e.Msg
		default:
			body.Code = -1
			body.Msg = e.Error()
		}
	} else {
		body.Code = 0
		body.Msg = "ok"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
