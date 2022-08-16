package api

import (
	feiraRoutes "github.com/fsvxavier/unico/internal/routes/feira"
	healthcheckRoutes "github.com/fsvxavier/unico/internal/routes/healthcheck"
	swaggerRoutes "github.com/fsvxavier/unico/internal/routes/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	root := app.Group("/", logger.New())

	// Setup the Banco Routes
	feiraRoutes.SetupFeirasRoutes(api)

	// Setup the Swagger Routes
	swaggerRoutes.SetupSwaggerRoutes(root)

	// Setup the HelthCheck Routes
	healthcheckRoutes.SetupHealthcheckRoutes(root)
}
