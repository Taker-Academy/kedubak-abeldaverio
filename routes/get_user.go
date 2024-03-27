package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func get_me_response_struct(user database_user) get_me_response {
	var to_return get_me_response

	to_return.Ok = true
	to_return.Data.Email = user.Email
	to_return.Data.FirstName = user.FirstName
	to_return.Data.LastName = user.LastName
	return to_return
}

func get_user(database *mongo.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		token_str := ctx.Request.Header.Get("Authorization")[7:]
		token, err := jwt.Parse(token_str, func(t *jwt.Token) (interface{}, error) {
			return []byte("SILTEPLAITMARCHE"), nil
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "invalid token")
		}
		exp, err := token.Claims.GetExpirationTime()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		if exp.Time.Before(time.Now()) {
			ctx.JSON(http.StatusBadRequest, "token has expired")
		}
		var token_id primitive.ObjectID
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			token_id, err = primitive.ObjectIDFromHex(claims["id"].(string))
			fmt.Println(token_id)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, ok)
		}
		collection := database.Database("KeDuBack").Collection("users")
		filter := bson.M{"_id": token_id}
		var user database_user
		err = collection.FindOne(context.Background(), filter).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
		}
		response := get_me_response_struct(user)
		ctx.JSON(http.StatusOK, response)
	}
	return (fn)
}
