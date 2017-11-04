package rest

import (
	"encoding/json"

	"code-lib/gerror"
	encoding_err "code-lib/gerror/encoding"
)

type RestResponse struct {
	// 错误码
	// 0: OK
	ErrorCode int64 `json:"ErrorCode"`

	// 错误原因
	ErrorMsg string `json:"ErrorMsg,omitempty"`

	// 请求唯一ID
	SessionID string `json:"SessionID"`

	// 响应内容
	Data interface{} `json:"Data,omitempty"`
}

func CreateRestResponse(sessionid string) (resp RestResponse) {
	resp = RestResponse{
		ErrorCode: 0,
		SessionID: sessionid,
	}

	return
}

func NewRestResponse(sessionid string) (resp *RestResponse) {
	r := CreateRestResponse(sessionid)
	resp = &r
	return
}

func (this *RestResponse) SetGError(gerr gerror.Error) *RestResponse {
	if gerr.IsNil() {
		this.Clear()
		return this
	}

	this.ErrorCode = gerr.ErrCode()
	this.ErrorMsg = gerr.Error()
	return this
}

func (this *RestResponse) Clear() {
	this.Data = nil
	this.ErrorCode = 0
	this.ErrorMsg = ""
}

func (this *RestResponse) SetData(data interface{}) {
	this.Data = data
}

func (this *RestResponse) IsOk() bool {
	if this == nil {
		return false
	}

	return this.ErrorCode == 0
}

func (this *RestResponse) Marshal2JSON() (data []byte, gerr *gerror.GError) {
	var err error
	data, err = json.Marshal(this)
	if err != nil {
		gerr = encoding_err.ErrJSONMarshal(err)
		return
	}

	return
}

func (this *RestResponse) Unmarshal2JSON(data []byte) (gerr *gerror.GError) {
	var err error
	err = json.Unmarshal(data, this)
	if err != nil {
		gerr = encoding_err.ErrJSONUnmarshal(err)
		return
	}

	return
}
