package main

import (
	"fmt"
	"github.com/Golang-Personal-Projects/Go-Projects/10-Go-Fiber-CRM/database"
	"github.com/Golang-Personal-Projects/Go-Projects/10-Go-Fiber-CRM/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

// setupRoutes Function
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

// initDB Function
func initDB() {
	var err error

	database.DBConn, err = gorm.Open("postgres", "host=berry.db.elephantsql.com port=5432 user=ncnezuuv dbname=ncnezuuv password=u5chPXwCwDT4EP8SYEimznP-8G9QQWyM")

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
