package main

import (
	"fmt"
	"github.com/Golang-Personal-Projects/Go-Projects/10-Go-Fiber-CRM/database"
	"github.com/Golang-Personal-Projects/Go-Projects/10-Go-Fiber-CRM/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

// setupRoutes Function
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

var (
	DbPassword = os.Getenv("DB_PASSWORD")
	DbUser     = os.Getenv("DB_USER")
	DbName     = os.Getenv("DB_NAME")
	DbHost     = os.Getenv("DB_HOST")
	DbPort     = os.Getenv("DB_PORT")
)

// initDB Function
func initDB() {
	var err error

	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	database.DBConn, err = gorm.Open("postgres", args)

	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
}

// main Function
func main() {
	// Create a new Fiber Instance
	app := fiber.New()

	// initialize database
	initDB()

	// Setup Routes
	setupRoutes(app)

	// start the server
	err := app.Listen(3000)
	if err != nil {
		log.Fatalf("unable to start the server %v", err)
	}

	defer database.DBConn.Close()

}
