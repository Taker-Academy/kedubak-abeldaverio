package routes

import (
	"KeDuBack/token"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func get_remove_response(user database_user) remove_response {
	var to_return remove_response
	to_return.Ok = true
	to_return.Data.Email = user.Email
	to_return.Data.FirstName = user.FirstName
	to_return.Data.LastName = user.LastName
	to_return.Data.Removed = true
	return to_return
}

func remove_user(database *mongo.Client) gin.HandlerFunc {
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
		_, err = collection.DeleteOne(context.Background(), filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		repsonse := get_remove_response(user)
		ctx.JSON(http.StatusOK, repsonse)
	}
	return (fn)
}
