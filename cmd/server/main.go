package main

import (
	"log"

	"user-api/config"
	"user-api/db/sqlc"
	"user-api/internal/logger"
	"user-api/internal/middleware"
	"user-api/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = logger.Init()
	if err != nil {
		log.Fatal("Logger failed:", err)
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	queries := sqlc.New(db)

	app := fiber.New()

	app.Use(middleware.RequestMiddleware())

	routes.SetupRoutes(app, queries)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "User API is running",
		})
	})

	logger.Log.Info("Server running on port 8080")
	log.Fatal(app.Listen(":8080"))
}