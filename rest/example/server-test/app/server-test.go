package main

import (
	"fmt"

	"code-lib/rest/restserver/resthttp"

	"rest_example/server-test/restmux"
)

func main() {
	httpSvr := resthttp.NewHTTPServer("localhost:7777", restmux.HTTPMux)
	if gerr := httpSvr.Serve(); !gerr.IsNil() {
		fmt.Println(gerr.Error())
	}
}
