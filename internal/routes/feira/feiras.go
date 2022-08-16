package feiraRoutes

import (
	feiraHandler "github.com/fsvxavier/unico/internal/handlers/feira"
	"github.com/gofiber/fiber/v2"
)

func SetupFeirasRoutes(router fiber.Router) {

	feira := router.Group("/feiras")

	// Create a Feira
	feira.Post("/", feiraHandler.CreateFeira)

	// Read Feiras Pagination
	feira.Get("/p/:page", feiraHandler.GetFeirasPagination)

	// Read Feiras Search
	feira.Post("/search", feiraHandler.GetFeiraSearch)

	// Read one Feira
	feira.Get("/:id", feiraHandler.GetFeira)

	// Update one Feira
	feira.Put("/:id", feiraHandler.UpdateFeira)

	// Delete one Feira
	feira.Delete("/:id", feiraHandler.DeleteFeira)
}
