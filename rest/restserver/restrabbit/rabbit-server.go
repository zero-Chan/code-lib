package restrabbit

import (
	"code-lib/gerror"
	//	system_err "code-lib/gerror/system"
)

type RabbitServer struct {
	addr string
	mux  RabbitMux
}

func CreateRabbitServer(addr string, mux RabbitMux) RabbitServer {
	server := RabbitServer{
		addr: addr,
		mux:  mux,
	}

	return server
}

func NewRabbitServer(addr string, mux RabbitMux) *RabbitServer {
	server := CreateRabbitServer(addr, mux)
	return &server
}

func (this *RabbitServer) Serve() (gerr gerror.Error) {
	var (
	//		err error
	)

	//	err = http.ListenAndServe(this.addr, this)
	//	if err != nil {
	//		return system_err.ErrHTTP(err)
	//	}

	return
}

//func (this *RabbitServer) ServeHTTP(respw http.ResponseWriter, req *http.Request) {
//	builder, gerr := this.mux.FindBuilder(req)
//	if !gerr.IsNil() {
//		respw.WriteHeader(http.StatusNotFound)
//		respw.Write(gerr.ErrorBytes())
//		return
//	}

//	ctl := newHTTPController(this, builder)
//	ctl.ServeHTTP(respw, req)
//}
