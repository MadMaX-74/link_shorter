package jwt

import "github.com/golang-jwt/jwt/v5"

type JwtData struct {
	Email string `json:"email"`
}

type Jwt struct {
	Secret string
}

func NewJwt(secret string) *Jwt {
	return &Jwt{
		Secret: secret,
	}
}

func (j *Jwt) GenerateToken(data JwtData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
	})
	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *Jwt) ParseToken(token string) (bool, *JwtData) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	email := t.Claims.(jwt.MapClaims)["email"]
	return t.Valid, &JwtData{
		Email: email.(string),
	}
}
