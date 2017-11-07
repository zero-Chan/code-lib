package rest

type HandlerFunc func() *RestResponse

type Handler interface {
	Prepare() *RestResponse
	Handle() *RestResponse
	Finish() *RestResponse
	RestChan() *RestChannel
}
