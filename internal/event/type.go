package event

// Entity 事件实体，包含事件的所有元信息和加密数据
type Entity struct {
	Topic     string `json:"topic"`     // 事件主题
	Operation string `json:"operation"` // 事件操作类型
	Time      int64  `json:"time"`      // 事件时间戳
	Nonce     string `json:"nonce"`     // 随机字符串，防重放
	Signature string `json:"signature"` // 事件签名
	Data      string `json:"data"`      // 加密的事件数据
}
