package routes

import (
	"KeDuBack/token"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func get_me_posts(database *mongo.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		token_id, err, code := token.Verify_token(ctx)
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
			if new.UserId == token_id.Hex() {
				response.Data = append(response.Data, get_post_response(new))
			}
		}
		response.Ok = true
		ctx.JSON(http.StatusOK, response)
	}
	return (fn)
}
