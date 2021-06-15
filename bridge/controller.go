package bridge

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller Struct

type ServiceResponse struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func (r *ServiceResponse) WriteResponse(w http.ResponseWriter) {
	response, _ := json.Marshal(*r)
	fmt.Fprintln(w, string(response))
}
