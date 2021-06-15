package security

type TokenValidator struct {
	Validator       func(token string) (isValid bool, err error)
}

func (t *TokenValidator) SetValidator(validator func(token string) (isValid bool, err error)) {
	t.Validator = validator
}

func (t *TokenValidator) ValidationAgent(token string) (bool, error) {
	return t.Validator(token)
}