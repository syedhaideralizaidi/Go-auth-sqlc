package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-auth-sqlc/controller" // Ensure correct import path
	"go-auth-sqlc/database"   // Ensure correct import path
)

func main() {
	database.ConnectDB()
	defer database.Conn.Close(context.Background())

	r := gin.Default()

	r.POST("/users", controller.CreateUser)
	r.POST("/users/:id", controller.GetUser)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
