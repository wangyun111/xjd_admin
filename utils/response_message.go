package utils

type replayMsg struct {
	Code   int    `json:"code"`
	SubMsg string `json:"subMsg"`
}

type replayData struct {
	Data string `json:"data"`
}

//操作成功
func OperationSuccess(msg ...string) replayMsg {
	result := "操作成功"
	if len(msg) > 0 {
		result = msg[0]
	}
	return replayMsg{Code: 200, SubMsg: result}
}

//操作失败
func OperationFalse(msg ...string) replayMsg {
	result := "操作失败"
	if len(msg) > 0 {
		result = msg[0]
	}
	return replayMsg{Code: 304, SubMsg: result}
}

//参数为空
func ParamNull(msg ...string) replayMsg {
	result := "参数为空"
	if len(msg) > 0 {
		result = msg[0]
	}
	return replayMsg{Code: 404, SubMsg: result}
}

//参数为空
func ParamError(msg ...string) replayMsg {
	result := "参数为空"
	if len(msg) > 0 {
		result = msg[0]
	}
	return replayMsg{Code: 401, SubMsg: result}
}

//返回结果为空
func ReturnDataNull(msg ...string) replayMsg {
	result := "返回值结果为空。"
	if len(msg) > 0 {
		result = msg[0]
	}
	return replayMsg{Code: 400, SubMsg: result}
}

// 自定义返回对象
func GetReturnObject(str string) replayData {
	return replayData{Data: str}
}
