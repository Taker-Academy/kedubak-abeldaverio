package main

import (
	"KeDuBack/cors_handler"
	"KeDuBack/database"
	"KeDuBack/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database := database.Init_database()

	router := gin.Default()
	router.Use(cors_handler.Setup_Header())

	routes.Setup_routes(router, database)
	router.Run("0.0.0.0:8080")
}
