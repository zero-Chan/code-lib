package restrabbit

import (
	"github.com/streadway/amqp"
)

type AMQPConsumer struct {
	// 队列名
	Queue string

	// 消费者名，为空则自动生成一个唯一ID
	Consumer string

	// 为true时收到消息自动ack
	AutoAck bool

	// 独有队列，只有该消费者才能订阅该队列
	Exclusive bool

	// 为true时不等待rabbitmqServer确认请求，立即开始交付
	NoWait bool

	Args amqp.Table

	Builder RabbitBuilder
}

type RabbitMux interface {
	RangeConsumer() []AMQPConsumer
}
