package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	db "go-auth-sqlc/database/sqlc" // Ensure correct import path
	"log"
)

var Conn *pgx.Conn
var Queries *db.Queries

func ConnectDB() {
	var err error
	Conn, err = pgx.Connect(context.Background(), "postgres://root:secret@localhost:5432/simple_auth?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	Queries = db.New(Conn) // Initialize Queries with the connection
}
