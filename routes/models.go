package routes

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type new_user struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type database_new_user struct {
	CreatedAt  time.Time `json:"createdAt"`
	Email      string    `json:"email"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Password   string    `json:"password"`
	LastUpVote time.Time `json:"lastUpVote"`
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

type login_infos struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type database_user struct {
	Id         primitive.ObjectID `bson:"_id", json:"id"`
	CreatedAt  time.Time          `json:"createdAt"`
	Email      string             `json:"email"`
	FirstName  string             `json:"firstName"`
	LastName   string             `json:"lastName"`
	Password   string             `json:"password"`
	LastUpVote time.Time          `json:"lastUpVote"`
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
	Ok   bool `json:"ok"`
	Data user `json:"user"`
}
