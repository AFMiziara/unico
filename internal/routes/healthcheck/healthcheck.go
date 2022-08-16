package healthcheckRoutes

import (
	"github.com/gofiber/fiber/v2"
	healthcheckHandler "github.com/fsvxavier/unico/internal/handlers/healthcheck"
)

func SetupHealthcheckRoutes(router fiber.Router) {

	healthcheck := router.Group("/healthcheck")

	healthcheck.Get("/", healthcheckHandler.GetHealthcheck)

}
