package rabbitmq

import (
	"strconv"

	"github.com/streadway/amqp"
)

type RabbitClientConf struct {
	Host     string
	Port     int64
	UserName string
	Password string
	VHost    string
}

type RabbitNotifyConf struct {
	*RabbitClientConf

	PublisherInuse bool
	ConsumerInuse  bool

	// Publisher need
	Exchange string

	// Publisher need
	RoutingKey string

	// Consumer need
	QueueName string

	// 高级选项
	Advanced *Table

	// TODO
	// Banned user declare exchange, queue and bind routingkey.
	Kind string // direct , fanout , topic
}

func (this *RabbitClientConf) Addr() string {
	if len(this.Host) == 0 || this.Port <= 0 {
		return ""
	}

	return this.Host + ":" + strconv.FormatInt(this.Port, 10)
}

// amqp_URI = "amqp:// amqp_authority ["/" vhost] ["?" query]
// amqp_authority = [amqp_userinfo "@"] host [":" port]
// amqp_userinfo = username [":" password]
// username = *(unreserved / pct-encoded / sub-delims)
// password = *(unreserved / pct-encoded / sub-delims)
// vhost = segment
func (this *RabbitClientConf) String() string {
	amqpUri := amqp.URI{
		Scheme:   "amqp",
		Host:     this.Host,
		Port:     int(this.Port),
		Username: this.UserName,
		Password: this.Password,
		Vhost:    this.VHost,
	}

	return amqpUri.String()
}

type Table map[string]interface{}
