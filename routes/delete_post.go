package routes

import (
	"KeDuBack/token"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func delete_post(database *mongo.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		_, err, code := token.Verify_token(ctx)
		if err != nil {
			ctx.JSON(code, err)
			return
		}
		post_id := ctx.Param("id")
		if post_id == "" {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		id, err := primitive.ObjectIDFromHex(post_id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		filter := bson.M{"_id": id}
		collection := database.Database("KeDuBack").Collection("posts")
		delete_response, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		fmt.Println(delete_response)

	}
	return (fn)
}
