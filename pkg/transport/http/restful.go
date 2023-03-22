package http

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mimani68/policy-engine/core"
	"github.com/mimani68/policy-engine/data"
)

func InitHttpServer(port int, policyInstance core.IPolicy) {
	ADDRESS := fmt.Sprintf(":%d", port)
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Restful policy manager is up and running",
		})
	})

	app.Post("/validate", func(c *fiber.Ctx) error {
		r := new(data.AccessRequestDto)
		if err := c.BodyParser(r); err != nil {
			return err
		}

		checkResult := policyInstance.CheckPolicy(r.Realm, "name", r.Name)

		return c.JSON(map[string]interface{}{
			"message": "Your policy validated",
			"allowed": checkResult,
		})
	})

	log.Fatal(app.Listen(ADDRESS))
}
