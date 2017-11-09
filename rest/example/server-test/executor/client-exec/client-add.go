package client_exec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	encoding_err "code-lib/gerror/encoding"
	system_err "code-lib/gerror/system"

	"code-lib/rest"
	"code-lib/rest/restserver/resthttp"

	prot "rest_example/server-test/prot/exec-prot/client-exec"
)

type ClientAddBuilder struct {
}

func NewClientAddBuilder() *ClientAddBuilder {
	builder := &ClientAddBuilder{}

	return builder
}

func (this *ClientAddBuilder) BuildExectorFromHTTP(httpch *resthttp.HTTPChannel) (rest.Executor, *rest.RestResponse) {
	var (
		rawMessage []byte
		request    = &prot.ClientAddRequest{}
		resp       = rest.NewRestResponse(httpch.RestChan().ID)
	)

	rawMessage, err := ioutil.ReadAll(httpch.Request.Body)
	if err != nil {
		resp.SetGError(system_err.ErrIORead("HTTP", err))
		return nil, resp
	}

	err = json.Unmarshal(rawMessage, request)
	if err != nil {
		resp.SetGError(encoding_err.ErrJSONUnmarshal(err))
		return nil, resp
	}

	exec := NewClientAddExec(httpch.RestChan(), request)
	return exec, resp
}

type ClientAddExec struct {
	restChannel  *rest.RestChannel
	requestBody  *prot.ClientAddRequest
	responseBody *prot.ClientAddResponse
}

func NewClientAddExec(channel *rest.RestChannel, req *prot.ClientAddRequest) *ClientAddExec {
	exec := &ClientAddExec{
		restChannel: channel,
		requestBody: req,
	}

	return exec
}

func (this *ClientAddExec) Prepare() *rest.RestResponse {
	var (
		resp = rest.NewRestResponse(this.RestChan().ID)
	)

	return resp
}

func (this *ClientAddExec) Exec() *rest.RestResponse {
	var (
		resp = rest.NewRestResponse(this.RestChan().ID)
	)

	fmt.Printf("Client[%s] want to add data: %+v\n", this.RestChan().ID, this.requestBody)

	return resp
}

func (this *ClientAddExec) Finish() *rest.RestResponse {
	var (
		resp = rest.NewRestResponse(this.RestChan().ID)
	)

	resp.SetData(this.responseBody)

	return resp
}

func (this *ClientAddExec) RestChan() *rest.RestChannel {
	return this.restChannel
}
