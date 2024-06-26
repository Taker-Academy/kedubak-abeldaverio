package routes

import (
	"KeDuBack/token"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
		token_id, err, code := token.Verify_token(ctx)
		if err != nil {
			ctx.JSON(code, err)
			return
		}
		collection := database.Database("KeDuBack").Collection("users")
		filter := bson.M{"_id": token_id}
		var user database_user
		err = collection.FindOne(context.Background(), filter).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			return
		}
		response := get_me_response_struct(user)
		ctx.JSON(http.StatusOK, response)
	}
	return (fn)
}
