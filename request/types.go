package request

import (
	"encoding/base64"
	"strings"
)

type Header struct {
	Operate      string `json:"operate"`
	SessionToken string `json:"session_token"`
}

type Security struct {
	Token          string   `json:"token"`
	BasicAuth      string   `json:"basic_auth"`
	Scope          string   `json:"scope"`
	GrantType      string   `json:"grant_type"`
	basicAuthSplit []string `json:"-"`
}

func (s *Security) GetToken() string     { return s.Token }
func (s *Security) GetBasicAuth() string { return s.BasicAuth }
func (s *Security) encodeBasicAuth() *[]string {
	if nil != s.basicAuthSplit && len(s.basicAuthSplit) == 2 {
		return &s.basicAuthSplit
	}

	basicToken := s.BasicAuth[6:len(s.BasicAuth)]
	decodeToken, _ := base64.StdEncoding.DecodeString(basicToken)
	s.basicAuthSplit = strings.Split(string(decodeToken), ":")
	if len(s.basicAuthSplit) > 1 {
		return &s.basicAuthSplit
	}
	return nil
}
func (s *Security) GetBasicAuthUsername() string {
	basicAuth := s.encodeBasicAuth()
	return (*basicAuth)[0]
}

func (s *Security) GetBasicAuthPassword() string {
	basicAuth := s.encodeBasicAuth()
	return (*basicAuth)[1]
}

type Request struct {
	Header   Header      `json:"header"`
	Security Security    `json:"security"`
	Body     interface{} `json:"body"`
}
