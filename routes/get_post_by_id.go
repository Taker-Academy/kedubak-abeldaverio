package routes

import (
	"KeDuBack/token"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func get_new_post_response(post_id string, post database_new_post) new_post_response {
	var to_return new_post_response

	to_return.Ok = true
	to_return.Data.Id = post_id
	to_return.Data.Comments = []comments_response{}
	for i := 0; i < len(post.Comments); i++ {
		var new comments_response
		new.Content = post.Comments[i].Content
		new.FirstName = post.Comments[i].FirstName
		new.Id = post.Comments[i].Id
		new.CreatedAt = post.Comments[i].CreatedAt
		to_return.Data.Comments = append(to_return.Data.Comments, new)
	}
	to_return.Data.Content = post.Content
	to_return.Data.CreatedAt = post.CreatedAt
	to_return.Data.FirstName = post.FirstName
	to_return.Data.Title = post.Title
	to_return.Data.UpVotes = post.UpVotes
	to_return.Data.UserId = post.UserId
	return to_return
}

func get_post_by_id(database *mongo.Client) gin.HandlerFunc {
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
		var post database_new_post
		err = collection.FindOne(context.Background(), filter).Decode(&post)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		response := get_new_post_response(post_id, post)
		ctx.JSON(http.StatusOK, response)

	}
	return (fn)
}
