package rest

type Processor interface {
	Prepare() *RestResponse
	Finish() *RestResponse
}

type HandlerFunc func() *RestResponse

type Handler interface {
	Processor
	Handle() *RestResponse
}
