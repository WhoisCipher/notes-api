package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/WhoisCipher/notes-api/internal/handlers"
	"github.com/WhoisCipher/notes-api/internal/middleware"
	"github.com/WhoisCipher/notes-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Print("CouldNot read .env", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbUser == "" || dbPass == "" || dbPort == "" || dbName == "" {
		log.Fatal("Environment variables not set")
	}

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("CouldNot Connect to Database: ", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.Note{}); err != nil {
		log.Fatal("Migration Failed: ", err)
	}

	app := fiber.New()

	app.Post("/signup", handlers.Signup(db))
	app.Post("/login", handlers.Login(db))

	api := app.Group("/api", middleware.Authentication())

	api.Post("/notes", handlers.CreateNote(db))
    api.Get("/notes", handlers.GetNotes(db))
    api.Put("/notes/:id", handlers.UpdateNotes(db))
    api.Delete("/notes/:id", handlers.DeleteNote(db))

	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Couldnot start server: ", err)
	}
}
