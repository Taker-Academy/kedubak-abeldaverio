package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup_routes(router *gin.Engine, database *mongo.Client) {
	router.POST("/auth/register", create_user(database))
	router.POST("/auth/login", login(database))
	router.GET("/user/me", get_user(database))
	router.PUT("/user/edit", edit_user(database))
	router.DELETE("/user/remove", remove_user(database))
	router.GET("/post", get_posts(database))
	router.POST("/post", create_post(database))
	router.GET("/post/me", get_me_posts(database))
	router.GET("/post/:id", get_post_by_id(database))
	router.DELETE("/post/:id", delete_post(database))
	router.POST("/post/vote/:id", upvote(database))
	router.POST("/comment/:id", post_comment(database))
}
