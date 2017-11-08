package resthttp

import (
	"code-lib/rest"
)

type HTTPBuilder interface {
	BuildExectorFromHTTP(httpch *HTTPChannel) (exec rest.Executor, resp *rest.RestResponse)
}
