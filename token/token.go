package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Generate_token(ID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString([]byte("SILTEPLAITMARCHE"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
