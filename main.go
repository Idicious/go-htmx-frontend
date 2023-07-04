package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"thedcsherman/htmx/internal/profiles"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	r := gin.Default()

	r.Static("/public", "./web/public")
	r.LoadHTMLGlob("./web/templates/*")

	profile_service := profiles.NewProfileService(conn)

	profiles.CreateProfileRoutes(r, profile_service)

	// Start the server
	http.ListenAndServe(":8080", r)
}
