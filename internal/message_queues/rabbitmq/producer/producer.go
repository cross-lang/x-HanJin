// Package producer 提供 RabbitMQ 生产者功能，
// 负责连接队列、声明队列和发送消息。
package producer

import (
	"fmt"

	"github.com/streadway/amqp"
	"go.uber.org/zap"

	"x-HanJin/internal/config"
	"x-HanJin/pkg/log"
)

// Producer RabbitMQ 生产者
type Producer struct {
	conn    *amqp.Connection // AMQP 连接
	channel *amqp.Channel    // AMQP 通道
	queue   amqp.Queue       // 队列
}

// NewProducer 创建 RabbitMQ 生产者。
// 参数 queueName 为目标队列名称。
func NewProducer(queueName string) *Producer {
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		config.Cfg.RabbitMQUser, config.Cfg.RabbitMQPassword,
		config.Cfg.RabbitMQHost, config.Cfg.RabbitMQPort)

	log.Info("RabbitMQ 生产者正在连接",
		zap.String("host", config.Cfg.RabbitMQHost),
		zap.Int("port", config.Cfg.RabbitMQPort),
	)

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatal("RabbitMQ 生产者连接失败", zap.Error(err))
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("RabbitMQ 生产者创建通道失败", zap.Error(err))
	}

	// 声明队列
	q, err := ch.QueueDeclare(
		queueName, // 队列名称
		false,     // 是否持久化
		false,     // 是否自动删除
		false,     // 是否为排他队列
		false,     // 是否阻塞
		nil,       // 额外参数
	)
	if err != nil {
		log.Fatal("RabbitMQ 生产者声明队列失败", zap.Error(err))
	}

	log.Info("RabbitMQ 生产者已启动", zap.String("queue", queueName))
	return &Producer{
		conn:    conn,
		channel: ch,
		queue:   q,
	}
}

// Publish 发布消息到队列
func (p *Producer) Publish(message string) error {
	err := p.channel.Publish(
		"",           // exchange
		p.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Error("RabbitMQ 消息发送失败", zap.Error(err))
		return err
	}

	log.Info("RabbitMQ 消息已发送", zap.String("message", message))
	return nil
}

// Close 关闭生产者的通道和连接
func (p *Producer) Close() {
	if err := p.channel.Close(); err != nil {
		log.Error("关闭 RabbitMQ 生产者通道失败", zap.Error(err))
	}
	if err := p.conn.Close(); err != nil {
		log.Error("关闭 RabbitMQ 生产者连接失败", zap.Error(err))
	}
}
