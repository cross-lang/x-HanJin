package event

import (
	"context"
	"crypto/hmac"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"x-HanJin/internal/constants"
	"x-HanJin/pkg/log"
	"x-HanJin/pkg/utils"
)

// handlerFunc 事件处理函数类型
type handlerFunc func(ctx context.Context, data *Entity, msg string) error

// Process 事件处理程序，管理事件处理器注册和分发
type Process struct {
	handlerMap map[string]handlerFunc // 处理器注册表，key 为 "Topic.Operation"
	ak         string                 // 接入标识
	sk         string                 // 接入密钥
}

// NewProcess 创建事件处理程序实例
func NewProcess(ak, sk string) *Process {
	return &Process{
		handlerMap: make(map[string]handlerFunc),
		ak:         ak,
		sk:         sk,
	}
}

// RegisterHandler 注册事件处理器。
// name 格式为 "Topic.Operation"，如 "order.created"。
func (p *Process) RegisterHandler(name string, handler handlerFunc) {
	p.handlerMap[name] = handler
}

// Handle 处理事件：验签 → 解密 → 分发到对应处理器
func (p *Process) Handle(ctx context.Context, data *Entity) error {
	// 验证签名
	if !p.VerifySignature(data) {
		return fmt.Errorf("event signature verification failed")
	}

	// 解密事件数据
	cipher := utils.CalcMd5(p.sk)
	decryptedData, err := utils.Decrypt(data.Data, cipher, data.Nonce, constants.ModeCBC, constants.PKCS7Padding, constants.EncodingBase64)
	if err != nil {
		return fmt.Errorf("event data decryption failed: %w", err)
	}

	key := data.Topic + "." + data.Operation
	log.WithContext(ctx).Debug("事件处理",
		zap.String("topic", data.Topic),
		zap.String("operation", data.Operation),
	)

	// 查找并执行处理器
	fn, ok := p.handlerMap[key]
	if !ok {
		return errors.New("no handler registered for event: " + key)
	}

	return fn(ctx, data, decryptedData)
}

// VerifySignature 验证事件签名的合法性。
// 使用 HMAC-SHA256 进行签名比较，防止时序攻击。
func (p *Process) VerifySignature(data *Entity) bool {
	message := fmt.Sprintf("%s:%s:%s:%d:%s", p.ak, data.Topic, data.Nonce, data.Time, data.Data)
	expectedMAC := utils.CalcHMACSha256(message, p.sk)
	// 使用 hmac.Equal 进行时间安全的比较
	return hmac.Equal([]byte(expectedMAC), []byte(data.Signature))
}
