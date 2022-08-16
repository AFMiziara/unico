package swaggerRoutes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SetupSwaggerRoutes(router fiber.Router) {

	healthcheck := router.Group("/swagger")

	// Swagger
	healthcheck.Get("/*", swagger.HandlerDefault)

}
