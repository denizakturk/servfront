package kernel

import (
	"github.com/denizakturk/servfront/config"
	"github.com/denizakturk/servfront/security"
	"net/http"
)

type Config struct {
	Securing *config.Securing
	Crypt    *config.Crypt
	Router   *config.RouterHolderCluster
}

type Security struct {
	TokenValidator *security.TokenValidator
	Crypt          *security.Crypt
}

type Request struct {
	Token string
}

type Service struct {
	Config   *Config
	Security *Security
	Request  *Request
}

func (k *Service) Init() {
	k.Config = &Config{}
	k.Config.Securing = &config.Securing{}
	k.Config.Crypt = &config.Crypt{}
	k.Config.Router = &config.RouterHolderCluster{}

	k.Security = &Security{}
	k.Security.Crypt = &security.Crypt{}
	k.Security.TokenValidator = &security.TokenValidator{}

	k.Request = &Request{}
}

func (k *Service) SetCryptKeys(key, iv string) {
	k.Config.Crypt.Key = key
	k.Config.Crypt.IV = iv
}

func (k *Service) AddRoute(clusterName string, router *config.Router) {
	k.Config.Router.AddRouterToCluster(clusterName, router)
}

func (k *Service) SetTokenValidator(validator func(token string) (isValid bool, err error)) {
	k.Security.TokenValidator.Validator = validator
}

func (k *Service) TokenCatcher(r *http.Request) {
	k.Request.Token = r.Header.Get(k.Config.Securing.TokenHeaderName)
}

var Holder = &Service{}
