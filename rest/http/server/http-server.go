package server

import (
	//	"fmt"
	"net/http"

	"code-lib/gerror"
	system_err "code-lib/gerror/system"
)

type RoutingRule struct {
	Path   string
	Method string
}

type HTTPServer struct {
	addr string
	mux  HTTPMux

	// map[path]rulefunc
	routKeys map[string]http.HandlerFunc

	processor HTTPProcessor
}

func CreateHTTPServer(addr string, mux HTTPMux) HTTPServer {
	server := HTTPServer{
		addr:     addr,
		mux:      mux,
		routKeys: make(map[string]http.HandlerFunc),
	}

	return server
}

func NewHTTPServer(addr string, mux HTTPMux) *HTTPServer {
	server := CreateHTTPServer(addr, mux)
	return &server
}

func (this *HTTPServer) SetProcessor(processor HTTPProcessor) {
	this.processor = processor
}

func (this *HTTPServer) Serve() (gerr *gerror.GError) {
	var (
		err error
	)

	err = http.ListenAndServe(this.addr, this)
	if err != nil {
		return system_err.ErrHTTP(err)
	}

	return
}

func (this *HTTPServer) ServeHTTP(respw http.ResponseWriter, req *http.Request) {
	hdl, gerr := this.mux.FindHandler(req)
	if !gerr.IsNil() {
		respw.WriteHeader(http.StatusNotFound)
		respw.Write(gerr.ErrorBytes())
		return
	}

	ctl := this.newController(hdl)
	ctl.ServeHTTP(respw, req)
}

func (this *HTTPServer) newController(handler HTTPHandler) *HTTPController {
	ctl := NewHTTPController()
	ctl.svr = this
	ctl.Handler = handler
	return ctl
}
