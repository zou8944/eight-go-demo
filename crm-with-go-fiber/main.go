package main

import (
	"crm-with-go-filer/lead"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("/leads", lead.GetLeads)
	app.Get("/leads/:id", lead.GetLead)
	app.Post("/leads", lead.CreateLead)
	app.Delete("/leads/:id", lead.DeleteLead)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	_ = app.Listen(8080)
}
