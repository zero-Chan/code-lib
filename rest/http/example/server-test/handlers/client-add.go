package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	encoding_err "code-lib/gerror/encoding"
	"code-lib/rest"
	"code-lib/rest/http/server"

	"server-test/prot/handlers-pms"
)

type ClientAdd struct {
	httpch       *server.HTTPChannel
	requestBody  prot.ClientAddRequest
	responseBody prot.ClientAddResponse
}

func CreateClientAdd() ClientAdd {
	cli := ClientAdd{}

	return cli
}

func NewClientAdd() *ClientAdd {
	cli := CreateClientAdd()
	return &cli
}

func (this *ClientAdd) InitFromHTTP(httpch *server.HTTPChannel) {
	this.httpch = httpch
}

// unmarshal request
func (this *ClientAdd) Prepare() *rest.RestResponse {
	var (
		err     error
		bodybuf []byte
		resp    = rest.NewRestResponse(this.httpch.ID)
	)

	bodybuf, _ = ioutil.ReadAll(this.httpch.Request.Body)
	err = json.Unmarshal(bodybuf, &this.requestBody)
	if err != nil {
		gerr := encoding_err.ErrJSONUnmarshal(err)
		return rest.NewRestResponse(this.httpch.ID).SetGError(gerr)
	}

	return resp
}

// handle
func (this *ClientAdd) Handle() *rest.RestResponse {
	var (
		resp = rest.NewRestResponse(this.httpch.ID)
	)

	fmt.Printf("Client[%s] want to add data: %+v\n", this.httpch.ID, this.requestBody)

	return resp
}

// set response
func (this *ClientAdd) Finish() *rest.RestResponse {
	resp := rest.NewRestResponse(this.httpch.ID)
	resp.SetData(this.responseBody)

	return resp
}
