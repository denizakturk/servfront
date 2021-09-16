package bridge

import (
	"net/http"
)

type UrlParams struct {
	Params map[string]string
}

type Struct interface {
	MiddleWare(w http.ResponseWriter, r *http.Request, urlParams *UrlParams)
	EndpointRunner(func() *ServiceResponse)
}

type ResponseDTO interface {
	WriteResponse(w http.ResponseWriter)
}

type Matcher string
type Bridge interface{ Struct }
