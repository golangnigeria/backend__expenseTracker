package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golangnigeria/expenseTracker/internals/database"
	"github.com/golangnigeria/expenseTracker/internals/repository"
	"github.com/golangnigeria/expenseTracker/internals/repository/dbrepo"
	"github.com/golangnigeria/expenseTracker/internals/routes"
	"github.com/joho/godotenv"
)

type Application struct {
	DB  repository.DatabaseRepository
	DSN string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// initialize app
	app := fiber.New()

	// instance of the application
	var myApp Application

	// Connect to database
	connectingToDB, err := database.ConnectToDB()
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	myApp.DB = &dbrepo.PostgresDBRepo{DB: connectingToDB}
	// Close database connection when the app is closed
	defer func() {
		if err := myApp.DB.Connection().Close(); err != nil {
			log.Println("Error closing database connection")
		}
	}()

	// Routes
	routes.Routes(app)

	err = app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Error starting server")
	}
}
