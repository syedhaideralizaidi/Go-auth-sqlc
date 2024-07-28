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
	r.GET("/users/:id", controller.GetUser)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
	r.GET("/users", controller.ListUsers)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
