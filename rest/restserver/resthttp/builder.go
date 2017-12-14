package resthttp

import (
	"code-lib/rest"
)

type HTTPBuilder interface {
	HTTPBuildExec(httpch *HTTPCmd, resp *rest.RestResponse) (exec rest.Executor)
}
