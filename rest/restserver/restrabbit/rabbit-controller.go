package restrabbit

import (
	"github.com/streadway/amqp"

	"code-lib/rest"
)

type RabbitController struct {
	svr     *RabbitServer
	builder RabbitBuilder
}

func createRabbitController(svr *RabbitServer, builder RabbitBuilder) RabbitController {
	hc := RabbitController{
		svr:     svr,
		builder: builder,
	}
	return hc
}

func newRabbitController(svr *RabbitServer, builder RabbitBuilder) *RabbitController {
	hc := createRabbitController(svr, builder)
	return &hc
}

func (this *RabbitController) ServeRabbit(delivery amqp.Delivery) {

	var (
		restResp *rest.RestResponse
		executor rest.Executor
	)

	// new channel
	rabbitch := NewRabbitChannel(&delivery)

	for {
		// exec handler
		executor, restResp = this.builder.BuildExectorFromRabbit(rabbitch)
		if !restResp.IsOk() {
			break
		}

		restResp = executor.Prepare()
		if !restResp.IsOk() {
			break
		}

		restResp = executor.Exec()
		if !restResp.IsOk() {
			break
		}

		restResp = executor.Finish()
		if !restResp.IsOk() {
			break
		}

		break
	}

	if !restResp.IsOk() {
		delivery.Reject(false)
		return
	}

	delivery.Ack(false)
}
