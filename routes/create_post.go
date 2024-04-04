package routes

import (
	"KeDuBack/token"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func get_post_to_insert(post new_post, user database_user) database_new_post {
	var to_return database_new_post

	to_return.Content = post.Content
	to_return.CreatedAt = time.Now()
	to_return.UserId = user.Id.Hex()
	to_return.FirstName = user.FirstName
	to_return.Title = post.Title
	to_return.UpVotes = []string{}
	to_return.Comments = []comments{}
	return to_return
}

func get_create_post_response(id *mongo.InsertOneResult, to_insert database_new_post) new_post_response {
	var to_return new_post_response
	to_return.Ok = true
	to_return.Data.Id = id.InsertedID.(primitive.ObjectID).Hex()
	to_return.Data.Comments = []comments_response{}
	to_return.Data.Content = to_insert.Content
	to_return.Data.CreatedAt = to_insert.CreatedAt
	to_return.Data.FirstName = to_insert.FirstName
	to_return.Data.Title = to_insert.Title
	to_return.Data.UpVotes = to_insert.UpVotes
	to_return.Data.UserId = to_insert.UserId
	return to_return
}

func create_post(database *mongo.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		token_id, err, code := token.Verify_token(ctx)
		if err != nil {
			ctx.JSON(code, err)
			return
		}
		json_data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			panic(err)
		}
		var tmp new_post
		json.Unmarshal(json_data, &tmp)
		decoder := json.NewDecoder(strings.NewReader(string(json_data)))
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&tmp)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
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
		to_insert := get_post_to_insert(tmp, user)
		post_collection := database.Database("KeDuBack").Collection("posts")
		id, err := post_collection.InsertOne(context.Background(), to_insert)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			return
		}
		response := get_create_post_response(id, to_insert)
		ctx.JSON(http.StatusCreated, response)
	}
	return (fn)
}
