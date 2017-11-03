package server

import (
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
}

func CreateHTTPServer(addr string) HTTPMuxServer {
	server := HTTPMuxServer{
		addr: addr,
		mux:  http.NewServeMux(),
	}

	return server
}

func NewHTTPServer(addr string) *HTTPMuxServer {
	server := CreateHTTPServer(addr)
	return &server
}

func (this *HTTPMuxServer) RegistHandler(handler HTTPHandler, rule RoutingRule) {
	routingFunc := func(respw http.ResponseWriter, req *http.Request) {
		switch {
		case rule.Path != req.URL.Path:
			respw.WriteHeader(http.StatusNotFound)
			return
		case rule.Method != req.Method:
			respw.WriteHeader(http.StatusNotFound)
			return
		}

		ctl := NewHTTPController()
		ctl.svr = this
		ctl.Handler = handler
		ctl.ServeHTTP(respw, req)
	}

	this.mux.HandleFunc(rule.Path, routingFunc)

}

func (this *HTTPMuxServer) Serve() (gerr *gerror.GError) {
	var (
		err error
	)

	err = http.ListenAndServe(this.addr, this.mux)
	if err != nil {
		return system_err.ErrHTTP(err)
	}

	return
}
