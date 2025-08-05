package api

const (
	SuccessMsg = "success"
	FailMsg    = "fail"

	SuccessCode = 10000
	FailCode    = 10001
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func CustomResult(code int, msg string, data interface{}) *Result {
	return &Result{Code: code, Msg: msg, Data: data}
}

// ==================================================

func Success() *Result {
	return CustomResult(SuccessCode, SuccessMsg, nil)
}

func SuccessWithCode(code int) *Result {
	return CustomResult(code, SuccessMsg, nil)
}

func SuccessWithMsg(msg string) *Result {
	return CustomResult(SuccessCode, msg, nil)
}

func SuccessWithData(data interface{}) *Result {
	return CustomResult(SuccessCode, SuccessMsg, data)
}

func SuccessWithCodeMsg(code int, msg string) *Result {
	return CustomResult(code, msg, nil)
}

func SuccessWithCodeData(code int, data interface{}) *Result {
	return CustomResult(code, SuccessMsg, data)
}

func SuccessWithMsgData(msg string, data interface{}) *Result {
	return CustomResult(SuccessCode, msg, data)
}

// ==================================================

func Fail() *Result {
	return CustomResult(FailCode, FailMsg, nil)
}

func FailWithCode(code int) *Result {
	return CustomResult(code, FailMsg, nil)
}

func FailWithMsg(msg string) *Result {
	return CustomResult(FailCode, msg, nil)
}

func FailWithData(data interface{}) *Result {
	return CustomResult(FailCode, FailMsg, data)
}

func FailWithCodeMsg(code int, msg string) *Result {
	return CustomResult(code, msg, nil)
}

func FailWithCodeData(code int, data interface{}) *Result {
	return CustomResult(code, FailMsg, data)
}

func FailWithMsgData(msg string, data interface{}) *Result {
	return CustomResult(FailCode, msg, data)
}
