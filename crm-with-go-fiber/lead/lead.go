package lead

import (
	"github.com/gofiber/fiber"
	"math/rand"
	"strconv"
)

type Lead struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

var leads []Lead

func GetLeads(c *fiber.Ctx) {
	_ = c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	for _, lead := range leads {
		if lead.ID == id {
			_ = c.JSON(lead)
		}
	}
}

func CreateLead(c *fiber.Ctx) {
	var lead Lead
	if err := c.BodyParser(&lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	lead.ID = strconv.Itoa(rand.Intn(100000))
	leads = append(leads, lead)
	_ = c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	for index, lead := range leads {
		if lead.ID == id {
			leads = append(leads[:index], leads[index+1:]...)
			_ = c.JSON(lead)
			break
		}
	}
}
