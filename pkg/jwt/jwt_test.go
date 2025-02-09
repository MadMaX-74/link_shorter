package jwt_test

import (
	"go_dev/pkg/jwt"
	"testing"
)

func TestJwtCreate(t *testing.T) {
	const email = "test@test.com"
	jwtService := jwt.NewJwt("80etOm7Kn4eDg6Wk35GfPg/IbAuzMdMoW7yC28BxMK4=")
	token, err := jwtService.GenerateToken(jwt.JwtData{
		Email: email,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
	isValid, data := jwtService.ParseToken(token)
	if !isValid {
		t.Error("Token is not valid")
	}
	if data.Email != email {
		t.Error("Email is not equal")
	}
}
