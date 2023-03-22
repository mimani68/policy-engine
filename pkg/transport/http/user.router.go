package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mimani68/policy-engine/data"
)

func (r *Router) userRouter(c *fiber.Ctx) error {
	request := new(data.UserDto)
	if err := c.BodyParser(request); err != nil {
		return err
	}

	checkNameResult := r.policy.CheckPolicy(request.Realm, "name", request.Name)
	checkOrganizationResult := r.policy.CheckPolicy(request.Realm, "organization", request.Organization)

	return c.JSON(map[string]interface{}{
		"message": "Your policy validated",
		"allowed": checkNameResult && checkOrganizationResult,
	})
}
