// Package rabbitmq 提供 RabbitMQ 消息队列的连接和初始化功能，
// 包含默认生产者和消费者的创建。
package rabbitmq

import (
	"x-HanJin/internal/config"
	"x-HanJin/internal/message_queues/rabbitmq/consumer"
	"x-HanJin/internal/message_queues/rabbitmq/producer"

	"x-HanJin/pkg/log"
	"go.uber.org/zap"
)

// RMQdp 默认 RabbitMQ 生产者实例（RabbitMQ Default Producer）
var RMQdp *producer.Producer

// InitRabbitMQ 初始化 RabbitMQ 默认生产者和消费者
func InitRabbitMQ() {
	// 创建默认队列生产者
	RMQdp = producer.NewProducer(config.Cfg.RabbitMQDefaultQueueName)

	// 在后台 goroutine 中启动默认队列消费者
	go func() {
		rMQdc := consumer.NewConsumer(config.Cfg.RabbitMQDefaultQueueName)
		if err := rMQdc.Consume(); err != nil {
			log.Fatal("默认消费者启动失败", zap.Error(err))
		}
	}()

	log.Info("RabbitMQ 初始化完成")
}
