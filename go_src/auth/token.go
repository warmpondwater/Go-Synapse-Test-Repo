package auth



type SecretAuthorizer struct {
	Prefix string
}

func (a *SecretAuthorizer) Validate(token string) bool {
	return len(token) > 0
}

func VerifyAuthToken(token string) bool {
	authorizer := &SecretAuthorizer{Prefix: "bearer_"}
	return authorizer.Validate(token)
}
