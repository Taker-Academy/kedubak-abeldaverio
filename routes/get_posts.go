package routes

import (
	"KeDuBack/token"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func get_post_response(post database_post) response_post {
	var to_return response_post

	to_return.Id = post.Id.Hex()
	to_return.Comments = []comments_response{}
	for i := 0; i < len(post.Comments); i++ {
		var new comments_response
		new.Content = post.Comments[i].Content
		new.FirstName = post.Comments[i].FirstName
		new.Id = post.Comments[i].Id
		new.CreatedAt = post.Comments[i].CreatedAt
		to_return.Comments = append(to_return.Comments, new)
	}
	to_return.Content = post.Content
	to_return.CreatedAt = post.CreatedAt
	to_return.FirstName = post.FirstName
	to_return.Title = post.Title
	to_return.UpVotes = post.UpVotes
	to_return.UserId = post.UserId
	return to_return
}

func get_posts(database *mongo.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		_, err, code := token.Verify_token(ctx)
		if err != nil {
			ctx.JSON(code, err)
			return
		}
		collection := database.Database("KeDuBack").Collection("posts")
		json_posts, err := collection.Find(context.Background(), bson.D{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		var response post_list
		response.Data = []response_post{}
		for json_posts.Next(context.TODO()) {
			var new database_post
			err = json_posts.Decode(&new)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			response.Data = append(response.Data, get_post_response(new))
		}
		response.Ok = true
		ctx.JSON(http.StatusOK, response)
	}
	return (fn)
}
