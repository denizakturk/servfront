package kernel

func (k *Service) Checker() {
	k.TokenChecker()
}

func (k *Service) TokenChecker() {
	if k.Config.Securing.TokenValidate {
		if k.Security.TokenValidator == nil {
			panic("Validator not found!")
		}

		if k.Request.Token == "" {
			panic("Token not found!")
		}

		if ok, err := k.Security.TokenValidator.ValidationAgent(k.Request.Token); !ok {
			panic(err)
		}
	}
}
