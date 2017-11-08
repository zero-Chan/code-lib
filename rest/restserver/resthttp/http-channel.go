package resthttp

import (
	"net/http"

	"code-lib/rest"
)

type HTTPChannel struct {
	*rest.RestChannel

	Request   *http.Request
	Responsew http.ResponseWriter
}

func CreateHTTPChannel(respw http.ResponseWriter, req *http.Request) HTTPChannel {
	ch := HTTPChannel{
		RestChannel: rest.NewRestChannel(),
		Request:     req,
		Responsew:   respw,
	}
	return ch
}

func NewHTTPChannel(respw http.ResponseWriter, req *http.Request) *HTTPChannel {
	ch := CreateHTTPChannel(respw, req)
	return &ch
}

func (this *HTTPChannel) RestChan() *rest.RestChannel {
	return this.RestChannel
}
