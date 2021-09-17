package bridge

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller Struct

type ServiceResponse struct {
	Error error       `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func (r *ServiceResponse) WriteResponse(w http.ResponseWriter) {
	if nil != r.Error {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
	response, _ := json.Marshal(*r)
	fmt.Fprintln(w, string(response))
	r.Error = nil
	r.Data = nil
}
