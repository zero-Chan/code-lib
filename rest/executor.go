package rest

type Executor interface {
	Prepare() *RestResponse
	Exec() *RestResponse
	Finish() *RestResponse
	RestChan() *RestChannel
}
