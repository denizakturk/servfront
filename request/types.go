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
	token          string   `json:"token"`
	basicAuth      string   `json:"basic_auth"`
	basicAuthSplit []string `json:"-"`
}

func (s *Security) GetToken() string     { return s.token }
func (s *Security) GetBasicAuth() string { return s.basicAuth }
func (s *Security) encodeBasicAuth() *[]string {
	if len(s.basicAuthSplit) == 2 {
		return &s.basicAuthSplit
	}
	s.basicAuthSplit = strings.SplitAfterN(base64.StdEncoding.EncodeToString([]byte(s.basicAuth)), ":", 2)
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
	header   Header      `json:"header"`
	security Security    `json:"security"`
	Body     interface{} `json:"body"`
}
