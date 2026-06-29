package constants

// BaseResponse 统一响应结构体，用于封装 API 接口的返回数据。
type BaseResponse struct {
	Code  int    `json:"code"`           // 业务状态码
	Msg   string `json:"msg"`            // 响应消息
	Debug string `json:"debug,omitempty"` // 调试信息（仅 debug 模式输出）
}
