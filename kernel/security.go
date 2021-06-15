package kernel

import "errors"

func (k *Service) Checker() error {
	return k.TokenChecker()
}

func (k *Service) TokenChecker() error {
	if k.Config.Securing.TokenValidate {
		if k.Security.TokenValidator.Validator == nil {
			return errors.New("Validator not found!")
		}

		if k.Request.Token == "" {
			return errors.New("Token not found!")
		}

		if ok, err := k.Security.TokenValidator.ValidationAgent(k.Request.Token); !ok {
			return err
		}
	}
	return nil
}
