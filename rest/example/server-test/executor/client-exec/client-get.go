package client_exec

import (
	"fmt"

	"code-lib/rest"
	"code-lib/rest/restserver/resthttp"

	prot "rest_example/server-test/prot/exec-prot/client-exec"
)

type ClientGetBuilder struct {
}

func NewClientGetBuilder() *ClientGetBuilder {
	builder := &ClientGetBuilder{}

	return builder
}

func (this *ClientGetBuilder) BuildExectorFromHTTP(httpch *resthttp.HTTPChannel) (rest.Executor, *rest.RestResponse) {
	var (
		request = &prot.ClientGetRequest{}
		resp    = rest.NewRestResponse(httpch.RestChan().ID)
	)

	exec := NewClientGetExec(httpch.RestChan(), request)
	return exec, resp
}

type ClientGetExec struct {
	restChannel  *rest.RestChannel
	requestBody  *prot.ClientGetRequest
	responseBody *prot.ClientGetResponse
}

func NewClientGetExec(channel *rest.RestChannel, req *prot.ClientGetRequest) *ClientGetExec {
	exec := &ClientGetExec{
		restChannel: channel,
		requestBody: req,
	}

	return exec
}

func (this *ClientGetExec) Prepare() *rest.RestResponse {
	var (
		resp = rest.NewRestResponse(this.RestChan().ID)
	)

	return resp
}

func (this *ClientGetExec) Exec() *rest.RestResponse {
	var (
		resp = rest.NewRestResponse(this.RestChan().ID)
	)

	fmt.Printf("Client[%s] want to get data\n", this.RestChan().ID)

	return resp
}

func (this *ClientGetExec) Finish() *rest.RestResponse {
	var (
		resp = rest.NewRestResponse(this.RestChan().ID)
	)

	resp.SetData(this.responseBody)

	return resp
}

func (this *ClientGetExec) RestChan() *rest.RestChannel {
	return this.restChannel
}
