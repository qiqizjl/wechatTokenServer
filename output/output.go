package output

import "encoding/json"

type outPut struct {
	Status int         `json:"status"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
}

func MakeNotFound() string {
	return MakeJson(NotFound, "NotFound", "")
}

func MakeReqParamsError() string {
	return MakeJson(Error, "ParamsError", "")
}

func MakeSuccess(data interface{}) string {
	return MakeJson(Success, "OK", data)
}

func MakeJson(status int, errMsg string, data interface{}) string {
	returnJson := &outPut{
		Status: status,
		ErrMsg: errMsg,
		Data:   data,
	}
	b, err := json.Marshal(returnJson)
	if err != nil {
		return ""
	}
	return string(b)
}
