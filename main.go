package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-auth-sqlc/controller" // Ensure correct import path
	"go-auth-sqlc/database"   // Ensure correct import path
	"go-auth-sqlc/middleware"
)

func main() {
	logger := middleware.SetupLogger()
	database.ConnectDB()
	defer database.Conn.Close(context.Background())

	r := gin.Default()
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/verify-email", controller.VerifyEmail)
	r.POST("/request-password-reset", controller.RequestPasswordReset)
	r.POST("/reset-password", controller.ResetPassword)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware(logger))
	{
		protected.POST("/users", controller.CreateUser)
		protected.GET("/users/:id", controller.GetUser)
		protected.PUT("/users/:id", controller.UpdateUser)
		protected.DELETE("/users/:id", controller.DeleteUser)
		protected.GET("/users", controller.ListUsers)
	}

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
