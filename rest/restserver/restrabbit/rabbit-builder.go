package restrabbit

import (
	"code-lib/rest"
)

type RabbitBuilder interface {
	BuildExectorFromRabbit(rabbitch *RabbitChannel) (exec rest.Executor, resp *rest.RestResponse)
}
