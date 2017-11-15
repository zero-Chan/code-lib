package mux

import (
	//	"fmt"

	//	"code-lib/gerror"
	//	system_err "code-lib/gerror/system"

	"code-lib/rest/restserver/restrabbit"
)

type RabbitMuxRule struct {
	QueueName    string `required:"true"`
	ConsumerName string `required:"false"`

	//	RoutingKey string `required:"true"`
	//	Path     string `required:"true"`
}

type RabbitMux struct {
	routKeys []struct {
		Rule    RabbitMuxRule
		Builder restrabbit.RabbitBuilder
	}
}

func CreateRabbitMux() RabbitMux {
	mux := RabbitMux{
	//		routKeys: make(map[string]restrabbit.RabbitBuilder),
	}

	return mux
}

func NewRabbitMux() *RabbitMux {
	mux := CreateRabbitMux()
	return &mux
}

func (this *RabbitMux) RegisterBuilder(builder restrabbit.RabbitBuilder, rule RabbitMuxRule) {

}

func (this RabbitMux) RangeConsumer() {

}
