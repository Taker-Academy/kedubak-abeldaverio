package routes

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type new_user struct {
	Email     string `bson:"email", json:"email"`
	Password  string `bson:"password", json:"password"`
	FirstName string `bson:"firstName", json:"firstName"`
	LastName  string `bson:"lastName", json:"lastName"`
}

type database_new_user struct {
	CreatedAt  time.Time `bson:"createdAt", json:"createdAt"`
	Email      string    `bson:"email", json:"email"`
	FirstName  string    `bson:"firstName", json:"firstName"`
	LastName   string    `bson:"lastName", json:"lastName"`
	Password   string    `bson:"password", json:"password"`
	LastUpVote time.Time `bson:"lastUpVote", json:"lastUpVote"`
}

type user struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type data struct {
	Token json.Token `json:"token"`
	User  user       `json:"user"`
}

type response struct {
	Ok   bool `json:"ok"`
	Data data `json:"data"`
}

type edit_response struct {
	Ok   bool `json:"ok"`
	Data user `json:"data"`
}

type login_infos struct {
	Email    string `bson:"email", json:"email"`
	Password string `bson:"password", json:"password"`
}

type database_user struct {
	Id         primitive.ObjectID `bson:"_id", json:"id"`
	CreatedAt  time.Time          `bson:"createdAt", json:"createdAt"`
	Email      string             `bson:"email", json:"email"`
	FirstName  string             `bson:"firstName", json:"firstName"`
	LastName   string             `bson:"lastName", json:"lastName"`
	Password   string             `bson:"password", json:"password"`
	LastUpVote time.Time          `bson:"lastUpvote", json:"lastUpVote"`
}

type token_struct struct {
	token string
}

type claims_struct struct {
	id  string
	exp time.Time
	jwt.RegisteredClaims
}

type get_me_response struct {
	Ok   bool `bson:"ok", json:"ok"`
	Data user `bson:"data", json:"user"`
}

type remove_data struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Removed   bool   `json:"removed"`
}

type remove_response struct {
	Ok   bool        `json:"ok"`
	Data remove_data `json:"user"`
}

type new_post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type comments struct {
	CreatedAt time.Time `bson:"createdAt", json:"createdAt"`
	Id        string    `bson:"id", json:"id"`
	FirstName string    `bson:"firstName", json:"firstName"`
	Content   string    `bson:"comment", json:"comment"`
}

type comments_response struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        string    `json:"id"`
	FirstName string    `json:"firstName"`
	Content   string    `json:"content"`
}

type database_new_post struct {
	CreatedAt time.Time  `bson:"createdAt", json:"createdAt"`
	UserId    string     `bson:"userId", json:"userId"`
	FirstName string     `bson:"firstName", json:"firstName"`
	Title     string     `bson:"title", json:"title"`
	Content   string     `bson:"content", json:"content"`
	Comments  []comments `bson:"comments", json:"comments"`
	UpVotes   []string   `bson:"upVotes", json:"upVotes"`
}

type new_post_response struct {
	Ok   bool          `json:"ok"`
	Data response_post `json:"data"`
}

type database_post struct {
	Id        primitive.ObjectID `bson:"_id", json:"id"`
	CreatedAt time.Time          `bson:"createdAt", json:"createdAt"`
	UserId    string             `bson:"userId", json:"userId"`
	FirstName string             `bson:"firstName", json:"firstName"`
	Title     string             `bson:"title", json:"title"`
	Content   string             `bson:"content", json:"content"`
	Comments  []comments         `bson:"comments", json:"comments"`
	UpVotes   []string           `bson:"upVotes", json:"upVotes"`
}

type response_post struct {
	Id        string              `json:"_id"`
	CreatedAt time.Time           `json:"createdAt"`
	UserId    string              `json:"userId"`
	FirstName string              `json:"firstName"`
	Title     string              `json:"title"`
	Content   string              `json:"content"`
	Comments  []comments_response `json:"comments"`
	UpVotes   []string            `json:"upVotes"`
}

type post_list struct {
	Ok   bool            `json:"ok"`
	Data []response_post `json:"data"`
}

type upvote_response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type new_comment_response_data struct {
	FirstName string    `json:"firstName"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type new_comment_response struct {
	Ok   bool                      `json:"ok"`
	Data new_comment_response_data `json:"data"`
}
