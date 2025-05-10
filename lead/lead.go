package lead

import (
	"github.com/Golang-Personal-Projects/Go-Projects/10-Go-Fiber-CRM/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLead(c *fiber.Ctx) {
	// get params from context
	id := c.Params("id")

	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	err := c.JSON(lead)
	if err != nil {
		log.Fatalf("unable to id %v", err)
		return
	}
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn

	var leads []Lead
	db.Find(&leads)
	err := c.JSON(leads)
	if err != nil {
		log.Fatalf("unable to get leads %v", err)
		return
	}
}

func DeleteLead(c *fiber.Ctx) {
	db := database.DBConn
	id := c.Params("id")

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("no lead for with ID")
		return
	}
	db.Delete(&lead, id)
	c.Status(200).Send("Delete Successful")
}

func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	err := c.JSON(lead)
	if err != nil {
		log.Fatalf("unable to add new lead %v", err)
		return
	}
}
