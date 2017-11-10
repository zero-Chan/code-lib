package restrabbit

import (
	"github.com/streadway/amqp"

	"code-lib/rest"
)

type RabbitChannel struct {
	*rest.RestChannel

	delivery *amqp.Delivery
}

func CreateRabbitChannel(delivery *amqp.Delivery) RabbitChannel {
	ch := RabbitChannel{
		RestChannel: rest.NewRestChannel(),
		delivery:    delivery,
	}
	return ch
}

func NewRabbitChannel(delivery *amqp.Delivery) *RabbitChannel {
	ch := CreateRabbitChannel(delivery)
	return &ch
}

func (this *RabbitChannel) RestChan() *rest.RestChannel {
	return this.RestChannel
}
