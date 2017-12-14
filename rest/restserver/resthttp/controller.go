package resthttp

import (
	"net/http"

	"code-lib/rest"
)

type HTTPController struct {
	svr     *HTTPServer
	builder HTTPBuilder
}

func createHTTPController(svr *HTTPServer, builder HTTPBuilder) HTTPController {
	hc := HTTPController{
		svr:     svr,
		builder: builder,
	}
	return hc
}

func newHTTPController(svr *HTTPServer, builder HTTPBuilder) *HTTPController {
	hc := createHTTPController(svr, builder)
	return &hc
}

func (this *HTTPController) ServeHTTP(respw http.ResponseWriter, req *http.Request) {
	var (
		executor rest.Executor
	)
	respw.WriteHeader(http.StatusOK)

	// new http cmd
	cmd := NewHTTPCmd(respw, req)
	restResp := rest.NewRestResponse(cmd.ID)

	for {
		// exec handler
		executor = this.builder.HTTPBuildExec(cmd, restResp)
		if !restResp.IsOk() {
			break
		}

		executor.Prepare(restResp)
		if !restResp.IsOk() {
			break
		}

		executor.Exec(restResp)
		if !restResp.IsOk() {
			break
		}

		executor.Finish(restResp)
		if !restResp.IsOk() {
			break
		}

		break
	}

	bytes, _ := restResp.Marshal()
	respw.Write(bytes)
}
