package lead

import (
	"github.com/gofiber/fiber"
	"github.com/greshma007/go-fiber-crm-basic/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"` //golang to understand json
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int `json:"phone"`
}

// c *fiber.Ctx - to access params from user

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")  // get from request
	db := database.DbConn // instance of database
	var lead Lead         // type struct lead
	db.Find(&lead, id)    // will find lead in db with specific id
	c.JSON(lead)          // send respond in JSON
}

func GetLeads(c *fiber.Ctx) {
	db := database.DbConn
	var leads []Lead // use slice here - to get all the leads
	db.Find(&leads)
	c.JSON(leads)
}

func NewLead(c *fiber.Ctx) {

	db := database.DbConn
	var lead Lead

	// BodyParser - to parse the body sent by user
	if err := c.BodyParser(&lead); err !=nil { // if any error in parsing - err
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DbConn
	var lead Lead
	db.First(&lead,id)
	if lead.Name == "" { // no lead found
		c.Status(500).Send("No lead found with ID")
		return
	}
	db.Delete(&lead)
	c.JSON("Lead successfully deleted")
}
