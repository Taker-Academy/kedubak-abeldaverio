package routes

import (
	"KeDuBack/token"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func get_new_comment_response(comment comments) new_comment_response {
	var to_return new_comment_response

	to_return.Ok = true
	to_return.Data.FirstName = comment.FirstName
	to_return.Data.Content = comment.Content
	to_return.Data.CreatedAt = comment.CreatedAt
	return to_return
}

func generate_comment(json_data []byte, user database_user) comments {
	var new_comment comments
	json.Unmarshal(json_data, &new_comment)
	new_comment.FirstName = user.FirstName
	new_comment.CreatedAt = time.Now()
	new_comment.Id = uuid.New().String()
	return new_comment
}

func post_comment(database *mongo.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		token_id, err, code := token.Verify_token(ctx)
		if err != nil {
			ctx.JSON(code, err)
			return
		}
		post_id := ctx.Param("id")
		if post_id == "" {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		primitive_post_id, err := primitive.ObjectIDFromHex(post_id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		user_collection := database.Database("KeDuBack").Collection("users")
		filter := bson.M{"_id": token_id}
		var user database_user
		err = user_collection.FindOne(context.Background(), filter).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			return
		}
		post_collection := database.Database("KeDuBack").Collection("posts")
		filter = bson.M{"_id": primitive_post_id}
		var post database_post
		err = post_collection.FindOne(context.Background(), filter).Decode(&post)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			return
		}
		json_data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		new_comment := generate_comment(json_data, user)
		post.Comments = append(post.Comments, new_comment)
		_, err = post_collection.ReplaceOne(context.Background(), filter, post)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		response := get_new_comment_response(new_comment)
		ctx.JSON(http.StatusOK, response)
	}
	return (fn)
}
