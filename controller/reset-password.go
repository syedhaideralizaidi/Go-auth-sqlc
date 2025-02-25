package controller

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"go-auth-sqlc/database"
	db "go-auth-sqlc/database/sqlc"
	"go-auth-sqlc/utils"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"time"
)

func generateResetToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func RequestPasswordReset(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch user by email
	_, err := database.Queries.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	token, err := generateResetToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	expiry := time.Now().Add(1 * time.Hour)

	requestResetToken := db.UpdateResetTokenParams{
		Email:            req.Email,
		ResetTokenExpiry: pgtype.Timestamptz{Time: expiry, Valid: true},
		ResetToken:       pgtype.Text{String: token, Valid: true},
	}

	err = database.Queries.UpdateResetToken(context.Background(), requestResetToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update reset token"})
		return
	}

	err = utils.SendResetPasswordEmail(req.Email, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send reset email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset email sent"})
}

func ResetPassword(c *gin.Context) {
	var req struct {
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	requestResetPassword := db.ResetPasswordParams{
		Password:   string(hashedPassword),
		ResetToken: pgtype.Text{String: req.Token, Valid: true},
	}

	err = database.Queries.ResetPassword(context.Background(), requestResetPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successful"})
}
