package mux

import (
	"net/http"

	"code-lib/rest/http/server"
)

type MuxRule struct {
	Path string
}

type HTTPMux struct {
	*http.ServeMux
}

func CreateHTTPPathMux() HTTPMux {
	hm := HTTPMux{
		mux: http.NewServeMux(),
	}

	return hm
}

func NewHTTPPathMux() *HTTPMux {
	hm := CreateHTTPPathMux()
	return &hm
}

func (this *HTTPMux) RegisterHandler(handler server.HTTPHandler, rule MuxRule) {

}
