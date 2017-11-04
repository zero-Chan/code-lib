package restmux

import (
	"net/http"

	"code-lib/rest/http/server/mux"

	"server-test/handlers"
)

var (
	HTTPMux *httpMux
)

type httpMux struct {
	*mux.MPHTTPMux
}

func init() {
	HTTPMux = &httpMux{
		MPHTTPMux: mux.NewHTTPPathMux(),
	}

	regist()
}

func regist() {
	HTTPMux.MPHTTPMux.RegisterHandler(
		handlers.NewClientAdd(),
		mux.MPHTTPMuxRule{
			Path:   "/client/add",
			Method: http.MethodPost,
		},
	)

	//	HTTPMux.MPHTTPMux.RegisterHandler(
	//		handlers.NewClientDel(),
	//		mux.MPHTTPMuxRule{
	//			Path:   "/client/del",
	//			Method: http.MethodPost,
	//		},
	//	)
}
