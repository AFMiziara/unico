package healthcheckHandler

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary HealthCheck
// @Description HealthCheck API
// @Success 200
// @Router /healthcheck [get]
func GetHealthcheck(c *fiber.Ctx) error {
	// Else return notes
	return c.JSON(fiber.Map{"status": "success", "message": "Healthcheck OK"})
}
