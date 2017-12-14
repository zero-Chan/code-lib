package resthttp

import (
	"net/http"

	"code-lib/rest"
)

type HTTPCmd struct {
	*rest.Cmd

	Request   *http.Request
	Responsew http.ResponseWriter
}

func CreateHTTPCmd(respw http.ResponseWriter, req *http.Request) HTTPCmd {
	ch := HTTPCmd{
		Cmd:       rest.NewCmd(),
		Request:   req,
		Responsew: respw,
	}
	return ch
}

func NewHTTPCmd(respw http.ResponseWriter, req *http.Request) *HTTPCmd {
	ch := CreateHTTPCmd(respw, req)
	return &ch
}

func (this *HTTPCmd) RestCmd() *rest.Cmd {
	return this.Cmd
}
