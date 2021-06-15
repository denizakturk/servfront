package config

type Securing struct {
	EncryptResponse bool
	TokenValidate   bool
	TokenHeaderName string
}

func (s *Securing) SetTokenHeaderName(tokenHeaderName string) {
	s.TokenHeaderName = tokenHeaderName
}

func (s *Securing) EnableEncryptResponse() {
	s.EncryptResponse = true
}

func (s *Securing) DisableEncryptResponse() {
	s.EncryptResponse = false
}

func (s *Securing) EnableTokenValidation() {
	s.TokenValidate = true
}

func (s *Securing) DisableTokenValidation() {
	s.TokenValidate = false
}
