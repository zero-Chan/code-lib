package server

import (
	"fmt"
	"net/http"

	"code-lib/gerror"
	system_err "code-lib/gerror/system"
)

type RoutingRule struct {
	Path   string
	Method string
}

type HTTPMuxServer struct {
	addr string
	mux  *http.ServeMux

	// map[path]rulefunc
	routKeys map[string]http.HandlerFunc

	processor HTTPProcessor
}

func CreateHTTPServer(addr string) HTTPMuxServer {
	server := HTTPMuxServer{
		addr:     addr,
		mux:      http.NewServeMux(),
		routKeys: make(map[string]http.HandlerFunc),
	}

	return server
}

func NewHTTPServer(addr string) *HTTPMuxServer {
	server := CreateHTTPServer(addr)
	return &server
}

func (this *HTTPMuxServer) SetProcessor(processor HTTPProcessor) {
	this.processor = processor
}

func (this *HTTPMuxServer) RegistHandler(handler HTTPHandler, rule RoutingRule) {
	rulefunc, ok := this.routKeys[rule.Path]
	if !ok {
		// stack bottom func
		rulefunc = func(respw http.ResponseWriter, req *http.Request) {
			gerr := this.filter(req, rule)
			if !gerr.IsNil() {
				respw.WriteHeader(http.StatusNotFound)
				respw.Write(gerr.ErrorBytes())
			}

			ctl := this.newController(handler)
			ctl.ServeHTTP(respw, req)
		}
		this.routKeys[rule.Path] = rulefunc
		return
	}

	// push to stack
	newRulefunc := func(respw http.ResponseWriter, req *http.Request) {
		// a path support multiple method
		if !this.filter(req, rule).IsNil() {
			rulefunc(respw, req)
			return
		}

		ctl := this.newController(handler)
		ctl.ServeHTTP(respw, req)
	}

	this.routKeys[rule.Path] = newRulefunc

}

func (this *HTTPMuxServer) Serve() (gerr *gerror.GError) {
	var (
		err error
	)

	this.regist()

	err = http.ListenAndServe(this.addr, this.mux)
	if err != nil {
		return system_err.ErrHTTP(err)
	}

	return
}

func (this *HTTPMuxServer) regist() {
	for path, rulefunc := range this.routKeys {
		this.mux.HandleFunc(path, rulefunc)
	}
}

func (this *HTTPMuxServer) filter(req *http.Request, rule RoutingRule) (gerr *gerror.GError) {
	var (
		err error
	)

	switch {
	case rule.Method != req.Method:
		err = fmt.Errorf("Unsupport HTTP method(%s)", req.Method)
		return system_err.ErrHTTPMuxFilter(err)
	case rule.Path != req.URL.Path:
		err = fmt.Errorf("unmatch HTTP path(%s)", req.URL.Path)
		return system_err.ErrHTTPMuxFilter(err)
	}

	return
}

func (this *HTTPMuxServer) newController(handler HTTPHandler) *HTTPController {
	ctl := NewHTTPController()
	ctl.svr = this
	ctl.Handler = handler
	return ctl
}
