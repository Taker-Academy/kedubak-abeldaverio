package routes

import (
	"KeDuBack/encrypt"
	"KeDuBack/token"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func get_login_response(user database_user) (response, error) {
	var to_return response
	var err error

	to_return.Ok = true
	to_return.Data.Token, err = token.Generate_token(user.Id.Hex())
	if err != nil {
		return response{}, err
	}
	to_return.Data.User.Email = user.Email
	to_return.Data.User.FirstName = user.FirstName
	to_return.Data.User.LastName = user.LastName
	return to_return, nil
}

func login(database *mongo.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		json_data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			panic(err)
		}
		var tmp login_infos
		json.Unmarshal(json_data, &tmp)
		decoder := json.NewDecoder(strings.NewReader(string(json_data)))
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&tmp)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		collection := database.Database("KeDuBack").Collection("users")
		filter := bson.M{"email": tmp.Email}
		var user database_user
		err = collection.FindOne(context.Background(), filter).Decode(&user)
		fmt.Println(user)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "email not found"})
		}
		hash_password := encrypt.Hash_sring(tmp.Password)
		if hash_password != user.Password {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "incorect password"})
		}
		to_return, err := get_login_response(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusCreated, to_return)
	}
	return (fn)
}
