package server

import (
	"code-lib/rest"
)

type HTTPProcessor interface {
	InitFromHTTP(httpch *HTTPChannel)
	rest.Processor
}

type HTTPHandler interface {
	InitFromHTTP(httpch *HTTPChannel)
	rest.Handler
}
