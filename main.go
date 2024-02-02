package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/greshma007/go-fiber-crm-basic/database" // tell go to import the package
	"github.com/greshma007/go-fiber-crm-basic/lead"     // tell go to import the package
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// app is instance of go fiber
func setupRoutes(app *fiber.App) {
	
	// GET ALL
	app.Get("/api/v1/lead",lead.GetLeads)
	
	// GET LEAD
	app.Get("/api/v1/lead/:id",lead.GetLead)
	
	// NEW LEAD
	app.Post("/api/v1/lead",lead.NewLead)
	
	// DELETE LEAD
	app.Delete("/api/v1/lead/:id",lead.DeleteLead)
}

func initDatabase() {
	var err error 
	database.DbConn, err = gorm.Open("sqlite3","leads.db") // driver - sqlite3 , db - leads
	if err != nil {
		panic("Failed to connect db")
	}
	fmt.Println("Connection opened to db")
	database.DbConn.AutoMigrate(&lead.Lead{}) // lead - go file , Lead - struct
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DbConn.Close() // to close conn at the end of main() exec
}
