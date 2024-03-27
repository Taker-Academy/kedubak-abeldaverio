package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup_routes(router *gin.Engine, database *mongo.Client) {
	router.POST("/auth/register", create_user(database))
	router.POST("auth/login", login(database))
	router.GET("user/me", get_user(database))
}
