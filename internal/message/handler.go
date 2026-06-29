package message

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"x-HanJin/pkg/log"
)

// handlerFunc 消息处理函数类型
type handlerFunc func(ctx context.Context, data *Entity, msg *Message) error

// Process 消息处理程序，管理消息处理器注册和分发
type Process struct {
	handlerMap map[string]handlerFunc // 处理器注册表，key 为 "Source.Key"
}

// NewProcess 创建消息处理程序实例
func NewProcess() *Process {
	return &Process{
		handlerMap: make(map[string]handlerFunc),
	}
}

// RegisterHandler 注册消息处理器。
// name 格式为 "Source.Key"，如 "notification.push"。
func (p *Process) RegisterHandler(name string, handler handlerFunc) {
	p.handlerMap[name] = handler
}

// Handle 处理消息：验证 → 记录日志 → 分发到对应处理器
func (p *Process) Handle(ctx context.Context, data *Entity) error {
	if len(data.Msgs) < 1 {
		return errors.New("message list is empty")
	}

	key := data.Source + "." + data.Key

	log.WithContext(ctx).Debug("消息处理",
		zap.String("source", data.Source),
		zap.String("key", data.Key),
		zap.Int64("msgId", data.MsgId),
		zap.Int("msgCount", len(data.Msgs)),
	)

	// 查找并执行处理器
	fn, ok := p.handlerMap[key]
	if !ok {
		return errors.New("no handler registered for message: " + key)
	}

	return fn(ctx, data, data.Msgs[0])
}
