package routes

import (
	"KeDuBack/encrypt"
	"KeDuBack/token"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func get_edit_user_response(ctx_user new_user) edit_response {
	var response edit_response
	response.Ok = true
	response.Data.Email = ctx_user.Email
	response.Data.FirstName = ctx_user.FirstName
	response.Data.LastName = ctx_user.LastName
	return response
}

func edit_user(database *mongo.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		token_id, err, code := token.Verify_token(ctx)
		if err != nil {
			ctx.JSON(code, err)
			return
		}
		collection := database.Database("KeDuBack").Collection("users")
		filter := bson.M{"_id": token_id}

		json_data, err := io.ReadAll(ctx.Request.Body)
		var ctx_user new_user
		json.Unmarshal(json_data, &ctx_user)
		update := bson.M{"email": ctx_user.Email,
			"password":  encrypt.Hash_sring(ctx_user.Password),
			"firstName": ctx_user.FirstName,
			"lastName":  ctx_user.LastName}
		_, err = collection.UpdateOne(context.Background(), filter, bson.M{"$set": update})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		response := get_edit_user_response(ctx_user)
		ctx.JSON(http.StatusOK, response)
	}
	return (fn)
}
