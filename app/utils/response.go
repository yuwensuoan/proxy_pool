package utils

import "proxy_pool/app/global"

// 返回前端的数据格式
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 返回成功信息
func (R Response) Success(data interface{}) Response {
	return Response {
		Data: data,
		Code: global.CodeSuccess,
		Msg:  global.GetMsgByCode(global.CodeSuccess),
	}
}

// 返回错误信息
func (R Response) Error(code int) Response {
	return Response{
		Data: []interface{}{},
		Code: code,
		Msg:  global.GetMsgByCode(code),
	}
}