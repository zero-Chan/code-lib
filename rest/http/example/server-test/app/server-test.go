package main

import (
	"fmt"

	"code-lib/rest/http/server"

	"server-test/restmux"
)

func main() {
	httpSvr := server.NewHTTPServer("localhost:7777", restmux.HTTPMux)
	if gerr := httpSvr.Serve(); !gerr.IsNil() {
		fmt.Println(gerr.Error())
	}
}
