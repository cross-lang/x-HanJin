// Package consumer 提供 RabbitMQ 消费者功能，
// 负责连接队列、声明队列和消费消息。
package consumer

import (
	"fmt"

	"github.com/streadway/amqp"
	"go.uber.org/zap"

	"x-HanJin/internal/config"
	"x-HanJin/pkg/log"
)

// Consumer RabbitMQ 消费者
type Consumer struct {
	conn    *amqp.Connection // AMQP 连接
	channel *amqp.Channel    // AMQP 通道
	queue   amqp.Queue       // 队列
}

// NewConsumer 创建 RabbitMQ 消费者。
// 参数 queueName 为要消费的队列名称。
func NewConsumer(queueName string) *Consumer {
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		config.Cfg.RabbitMQUser, config.Cfg.RabbitMQPassword,
		config.Cfg.RabbitMQHost, config.Cfg.RabbitMQPort)

	log.Info("RabbitMQ 消费者正在连接",
		zap.String("host", config.Cfg.RabbitMQHost),
		zap.Int("port", config.Cfg.RabbitMQPort),
	)

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatal("RabbitMQ 消费者连接失败", zap.Error(err))
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("RabbitMQ 消费者创建通道失败", zap.Error(err))
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
		log.Fatal("RabbitMQ 消费者声明队列失败", zap.Error(err))
	}

	log.Info("RabbitMQ 消费者已启动", zap.String("queue", queueName))
	return &Consumer{
		conn:    conn,
		channel: ch,
		queue:   q,
	}
}

// Consume 开始消费消息，此方法会阻塞当前 goroutine。
// 消费到的消息会在独立的 goroutine 中处理。
func (c *Consumer) Consume() error {
	defer func(conn *amqp.Connection) {
		if err := conn.Close(); err != nil {
			log.Error("关闭 RabbitMQ 消费者连接失败", zap.Error(err))
		}
	}(c.conn)
	defer func(channel *amqp.Channel) {
		if err := channel.Close(); err != nil {
			log.Error("关闭 RabbitMQ 消费者通道失败", zap.Error(err))
		}
	}(c.channel)

	msgs, err := c.channel.Consume(
		c.queue.Name, // 队列名称
		"",           // 消费者标记
		true,         // 自动确认
		false,        // 是否排他
		false,        // 是否为本地消费者
		false,        // 队列的其他属性
		nil,          // 额外参数
	)
	if err != nil {
		log.Error("RabbitMQ 消费者注册失败", zap.Error(err))
		return err
	}

	// 处理消息
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Info("RabbitMQ 消费者收到消息", zap.ByteString("body", d.Body))
			// TODO: 在这里添加消息处理逻辑
		}
	}()

	<-forever
	return nil
}
