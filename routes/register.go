package routes

import (
	"KeDuBack/encrypt"
	"KeDuBack/token"
	"encoding/json"
	"net/http"

	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func get_to_insert(user new_user) database_new_user {
	var to_return database_new_user

	to_return.Email = user.Email
	to_return.FirstName = user.FirstName
	to_return.LastName = user.LastName
	to_return.Password = encrypt.Hash_sring(user.Password)
	to_return.CreatedAt = time.Now()
	to_return.LastUpVote = time.Now().Add(-1 * time.Minute)
	return (to_return)
}

func get_response(user new_user, id string) (response, error) {
	var to_return response
	var err error

	to_return.Ok = true
	to_return.Data.Token, err = token.Generate_token(id)
	if err != nil {
		return response{}, err
	}
	to_return.Data.User.Email = user.Email
	to_return.Data.User.FirstName = user.FirstName
	to_return.Data.User.LastName = user.LastName
	return to_return, nil
}

func create_user(database *mongo.Client) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		json_data, err := io.ReadAll(context.Request.Body)
		if err != nil {
			panic(err)
		}
		var tmp new_user
		json.Unmarshal(json_data, &tmp)
		decoder := json.NewDecoder(strings.NewReader(string(json_data)))
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&tmp)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		collection := database.Database("KeDuBack").Collection("users")
		var to_insert database_new_user = get_to_insert(tmp)
		check, err := collection.CountDocuments(context, bson.M{"email": to_insert.Email})
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if check > 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "this email is already used"})
			return
		}
		db_response, err := collection.InsertOne(context, to_insert)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		to_return, err := get_response(tmp, db_response.InsertedID.(primitive.ObjectID).Hex())
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		context.JSON(http.StatusCreated, to_return)
	}
	return (fn)
}
