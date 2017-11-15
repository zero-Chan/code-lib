package restrabbit

import (
	"github.com/streadway/amqp"

	"code-lib/gerror"
	//	system_err "code-lib/gerror/system"
)

type RabbitServer struct {
	addr     amqp.URI
	cli      *amqp.Connection
	channels []struct {
		amqpChan *amqp.Channel
		Consumer AMQPConsumer
	}

	mux RabbitMux
}

func CreateRabbitServer(addr amqp.URI, mux RabbitMux) RabbitServer {
	server := RabbitServer{
		addr: addr,
		mux:  mux,
	}

	return server
}

func NewRabbitServer(addr amqp.URI, mux RabbitMux) *RabbitServer {
	server := CreateRabbitServer(addr, mux)
	return &server
}

func (this *RabbitServer) Serve() (gerr gerror.Error) {
	var (
		err error
	)

	this.cli, err = amqp.Dial(this.addr.String())
	if err != nil {
		// TODO
		return
	}

	for _, pms := range this.mux.RangeConsumer() {
		ch, err := this.cli.Channel()
		deliveries, err := ch.Consume(pms.Queue, pms.Consumer, false, pms.Exclusive, false, pms.NoWait, pms.Args)
		if err != nil {
			// TODO
			return
		}

		go this.serveConsume(deliveries, pms.Builder)
	}

	return
}

func (this *RabbitServer) serveRabbit() {

}

func (this *RabbitServer) serveConsume(deliveries <-chan amqp.Delivery, builder RabbitBuilder) {
	for delivery := range deliveries {
		ctl := newRabbitController(this, builder)
		ctl.ServeRabbit(delivery)
	}
}
