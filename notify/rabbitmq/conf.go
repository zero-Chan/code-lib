package rabbitmq

import (
	"net/url"
	"strconv"
)

type RabbitClientConf struct {
	Host     string
	Port     int64
	UserName string
	Password string
	VHost    string
}

type RabbitMQConf struct {
	*RabbitClientConf
	Exchange       string
	RoutingKey     string
	QueueName      string
	BindKind       string // direct, fanout, topic
	PublisherInuse bool
	ConsumerInuse  bool

	// 高级选项
	Advanced *Table
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
	amqpUri := url.URL{}

	if this.Addr() == "" {
		return ""
	}

	amqpUri.Scheme = "amqp"
	amqpUri.User = url.UserPassword(this.UserName, this.Password)
	amqpUri.Host = this.Addr()

	if this.VHost != "" {
		amqpUri.Path = this.VHost
	}

	return amqpUri.String()
}

type Table map[string]interface{}
