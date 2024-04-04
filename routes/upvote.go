package routes

import (
	"KeDuBack/token"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func check_user(post_collection *mongo.Collection, post_id string, user_id string) bool {
	var post database_post
	post_id_primitive, _ := primitive.ObjectIDFromHex(post_id)

	_ = post_collection.FindOne(context.Background(), bson.M{"_id": post_id_primitive}).Decode(&post)
	for i := 0; i < len(post.UpVotes); i++ {
		if post.UpVotes[i] == user_id {
			return true
		}
	}
	return false
}

func upvote(database *mongo.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		token_id, err, code := token.Verify_token(ctx)
		if err != nil {
			ctx.JSON(code, err)
			return
		}
		post_collection := database.Database("KeDuBack").Collection("posts")
		users_collection := database.Database("KeDuBack").Collection("users")
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
		post_filter := bson.M{"_id": id}
		if check_user(post_collection, post_id, token_id.Hex()) {
			ctx.JSON(http.StatusUnauthorized, "you can't upvote twice")
			return
		}
		post_update := bson.M{"upVotes": token_id.Hex()}
		_, err = post_collection.UpdateOne(context.Background(), post_filter, bson.M{"$push": post_update})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		filter := bson.M{"_id": token_id}
		update := bson.M{"lastUpVote": time.Now()}
		_, err = users_collection.UpdateOne(context.Background(), filter, bson.M{"$set": update})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		var response upvote_response
		response.Ok = true
		response.Message = "post upvoted"
		ctx.JSON(http.StatusOK, response)
	}
	return (fn)
}
