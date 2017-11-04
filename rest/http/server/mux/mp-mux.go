package mux

import (
	"fmt"
	"net/http"

	"code-lib/gerror"
	system_err "code-lib/gerror/system"

	"code-lib/rest/http/server"
)

type MPHTTPMuxRule struct {
	Method string `required:"true"`
	Path   string `required:"true"`
}

// MPHTTPMux
// M: Method must to be hit
// P: Path must to be hit
type MPHTTPMux struct {
	// map[path][method]rulefunc
	routKeys map[string]map[string]server.HTTPHandler
}

func CreateHTTPPathMux() MPHTTPMux {
	hm := MPHTTPMux{
		routKeys: make(map[string]map[string]server.HTTPHandler),
	}

	return hm
}

func NewHTTPPathMux() *MPHTTPMux {
	hm := CreateHTTPPathMux()
	return &hm
}

func (this *MPHTTPMux) RegisterHandler(handler server.HTTPHandler, rule MPHTTPMuxRule) {
	// check must be set
	var (
		path   = rule.Path
		method = rule.Method
	)

	switch {
	case method == "":
		return
	case path == "":
		return
	}

	pathMethods, ok := this.routKeys[path]
	if !ok {
		// map[method]rulefunc
		pathMethods = make(map[string]server.HTTPHandler)
	}

	_, ok = pathMethods[method]
	if ok {
		panic(fmt.Sprintf("http: multiple registrations for path(%s), method(%s)", path, method))
	}

	pathMethods[method] = handler

	this.routKeys[path] = pathMethods
}

// FindHandler
// 查找算法高效与否严重影响性能
func (this *MPHTTPMux) FindHandler(req *http.Request) (hdl server.HTTPHandler, gerr *gerror.GError) {
	// Hits algorithm
	method := req.Method
	path := req.URL.Path

	ownMethods, ok := this.routKeys[path]
	if !ok {
		gerr = system_err.ErrHTTPMuxFilter(fmt.Errorf("Can not find HTTP path(%s)", path))
		return
	}

	hdl, ok = ownMethods[method]
	if !ok {
		gerr = system_err.ErrHTTPMuxFilter(fmt.Errorf("Can not find HTTP method(%s) in path(%s)", method, path))
		return
	}

	return
}
