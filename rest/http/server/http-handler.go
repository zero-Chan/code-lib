package server

import (
	"code-lib/rest"
)

type HTTPHandler interface {
	InitFromHTTP(httpch *HTTPChannel) *rest.RestResponse
	rest.Handler
}
