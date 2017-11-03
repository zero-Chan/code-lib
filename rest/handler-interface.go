package rest

type Handler interface {
	Handle() *RestResponse
}
