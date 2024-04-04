package token

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Verify_token(ctx *gin.Context) (primitive.ObjectID, any, int) {
	token_str := ctx.Request.Header.Get("Authorization")[7:]
	token, err := jwt.Parse(token_str, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_STRING")), nil
	})
	if err != nil {
		return primitive.ObjectID{}, "invalid token", http.StatusBadRequest
	}
	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	if exp.Time.Before(time.Now()) {
		return primitive.ObjectID{}, "token has expired", http.StatusBadRequest
	}
	var token_id primitive.ObjectID
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		token_id, err = primitive.ObjectIDFromHex(claims["id"].(string))
		if err != nil {
			return primitive.ObjectID{}, err, http.StatusInternalServerError
		}
	} else {
		return primitive.ObjectID{}, ok, http.StatusInternalServerError
	}
	return token_id, nil, 0
}

func Generate_token(ID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_STRING")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
