package application

type JWTToken struct {
	token string
}

func NewJWTToken(token string) JWTToken {
	return JWTToken{token: token}
}

func (jwt JWTToken) Token() string {
	return jwt.token
}
