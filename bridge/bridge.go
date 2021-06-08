package bridge

import "net/http"

type Struct interface {
	MiddleWare(w http.ResponseWriter, r *http.Request)
	EndpointRunner(func() *ServiceResponse)
}

type ResponseDTO interface {
	WriteResponse(w http.ResponseWriter)
}

type Matcher string
type Bridge interface{ Struct }
