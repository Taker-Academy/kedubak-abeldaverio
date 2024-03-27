package main

import (
	"KeDuBack/database"
	"KeDuBack/routes"
	"KeDuBack/cors_handler"

	"github.com/gin-gonic/gin"
)

func main() {
	database := database.Init_database()

	router := gin.Default()
	router.Use(cors_handler.Setup_Header())

	routes.Setup_routes(router, database)
	router.Run("localhost:8080")
}
