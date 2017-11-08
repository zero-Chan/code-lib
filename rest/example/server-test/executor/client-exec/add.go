package client_exec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	encoding_err "code-lib/gerror/encoding"
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
		// TODO set gerror

	}

	err = json.Unmarshal(rawMessage, request)
	if err != nil {
		resp.SetGError(encoding_err.ErrJSONUnmarshal(err))
		return nil, resp
	}

	exec := NewClientAddHandler(httpch.RestChan(), request)
	return exec, resp
}

type ClientAddExec struct {
	// datas
	restChannel  *rest.RestChannel
	requestBody  *prot.ClientAddRequest
	responseBody *prot.ClientAddResponse
}

func NewClientAddHandler(channel *rest.RestChannel, req *prot.ClientAddRequest) *ClientAddExec {
	hdl := &ClientAddExec{
		restChannel: channel,
		requestBody: req,
	}

	return hdl
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
