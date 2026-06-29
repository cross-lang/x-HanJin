package message

// Entity 消息实体，包含消息来源、目标用户和消息内容
type Entity struct {
	Source  string     `json:"source"`   // 消息来源标识
	Key     string     `json:"key"`      // 消息类型标识
	ToUsers []string   `json:"to_users"` // 目标用户列表
	MsgTime int64      `json:"msg_time"` // 消息时间戳
	Msgs    []*Message `json:"msgs"`     // 消息内容列表
	Extra   string     `json:"extra"`    // 扩展信息
	MsgId   int64      `json:"msg_id"`   // 消息唯一 ID
}

// Message 单条消息内容
type Message struct {
	Terminals []int64 `json:"terminals"` // 推送终端列表
	Category  string  `json:"category"`  // 消息分类
	Template  string  `json:"template"`  // 消息模板标识
	Body      string  `json:"body"`      // 消息正文
	Version   int64   `json:"version"`   // 消息版本号
	Ext       string  `json:"ext"`       // 扩展字段
	Nopop     bool    `json:"nopop"`     // 是否静默消息（不弹出通知）
}
