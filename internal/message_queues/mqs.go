// Package message_queues 提供消息队列初始化的统一入口，
// 支持 RabbitMQ、RocketMQ 和 Kafka。
package message_queues

import (
	"x-HanJin/internal/message_queues/kafka"
	"x-HanJin/internal/message_queues/rabbitmq"
	"x-HanJin/internal/message_queues/rocketmq"

	"x-HanJin/pkg/log"
	"go.uber.org/zap"
)

// Init 初始化所有消息队列
func Init() {
	log.Info("开始初始化消息队列...")

	rabbitmq.InitRabbitMQ()
	rocketmq.InitRocketMQ()
	kafka.InitKafka()

	log.Info("消息队列初始化完成", zap.Int("queues", 3))
}
