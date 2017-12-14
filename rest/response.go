package rest

import (
	"code-lib/gerror"
	"code-lib/rest/proto"
)

type RestResponse struct {
	baseResponse

	// 错误码
	// 0: OK
	ErrCode int64 `json:"ErrCode"`

	// 错误原因
	ErrMsg string `json:"ErrMsg,omitempty"`

	// 请求唯一ID
	SessionID string `json:"SessionID"`

	// 响应内容
	Data interface{} `json:"Data,omitempty"`
}

func CreateRestResponse(sessionid string) (resp RestResponse) {
	resp = RestResponse{
		baseResponse: CreatebaseResponse(),
		ErrCode:      0,
		SessionID:    sessionid,
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

	this.ErrCode = gerr.ErrCode()
	this.ErrMsg = gerr.Error()
	return this
}

func (this *RestResponse) Clear() {
	this.Data = nil
	this.ErrCode = 0
	this.ErrMsg = ""
}

func (this *RestResponse) SetData(data interface{}) {
	this.Data = data
}

func (this *RestResponse) IsOk() bool {
	if this == nil {
		return false
	}

	return this.ErrCode == 0
}

type baseResponse struct {
	Proto proto.ProtoType `json:"-"`
}

func CreatebaseResponse() baseResponse {
	resp := baseResponse{
		// default use json proto
		Proto: proto.ProtoType_JSON,
	}

	return resp
}

func (this *baseResponse) Marshal() (data []byte, gerr gerror.Error) {
	return this.Proto.Marshal(this)
}

func (this *baseResponse) Unmarshal(data []byte) gerror.Error {
	return this.Proto.Unmarshal(data, this)
}
