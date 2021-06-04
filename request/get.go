package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetRequest(r *http.Request) *Request {
	var readBody []byte
	var err error
	request := &Request{}
	readBody, err = ioutil.ReadAll(r.Body)
	if nil != err {
		panic("Body read error: " + err.Error())
	}
	json.Unmarshal(readBody, request)

	if "" == request.Header.Operate {
		request.Header.Operate = r.RequestURI
	}

	if "" == request.Header.SessionToken {
		cookie, cookieErr := r.Cookie("session_id")
		if nil == cookieErr {
			request.Header.SessionToken = cookie.Value
		}
	}

	return request
}
