package bridge

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller Struct

type ServiceResponse struct {
	Error error       `json:"error"`
	Data  interface{} `json:"data"`
}

func (r *ServiceResponse) WriteResponse(w http.ResponseWriter) {
	response, _ := json.Marshal(*r)
	fmt.Fprintln(w, string(response))
}
