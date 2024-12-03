package main

import (
	"go-blog/database"
	"go-blog/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()
	auth := r.Group("/auth")
	routes.AuthRoutes(auth)
	r.Run(":3000")
}
