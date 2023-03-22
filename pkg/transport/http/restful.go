package http

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mimani68/policy-engine/core"
)

func InitHttpServer(port int, policyInstance core.IPolicy) {
	ADDRESS := fmt.Sprintf(":%d", port)
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Restful policy manager is up and running",
		})
	})

	router := NewRouter(policyInstance)
	app.Post("/validate/user", router.userRouter)

	log.Fatal(app.Listen(ADDRESS))
}
