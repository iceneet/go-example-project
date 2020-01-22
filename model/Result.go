package model

//返回结构
type Result struct{
	Code    int         `json:"code" example:"000"`
	Message string      `json:"message" example:"请求信息"`
	Data    interface{} `json:"data" `
}

func ResultInit(code int, message string, data interface{}) Result {
	return Result{
		Code:    code,
		Message: message,
		Data:    data,
	}
}