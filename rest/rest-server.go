package rest

import (
	"code-lib/gerror"
)

type RestServer interface {
	Serve() (gerr gerror.Error)
}
