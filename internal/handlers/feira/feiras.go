package feiraHandler

import (
	"fmt"

	"github.com/fsvxavier/unico/database"
	feiraInterfaces "github.com/fsvxavier/unico/internal/interfaces/feira"
	"github.com/fsvxavier/unico/internal/models"
	feiraRepository "github.com/fsvxavier/unico/internal/repositories/feira"
	feiraUsecases "github.com/fsvxavier/unico/internal/usecases/feira"
	"github.com/gofiber/fiber/v2"
)

var (
	Feiras           []models.FeiraLivre
	Feira            models.FeiraLivre
	SearchFeira      models.SearchFeira
	InUpFeiras       models.InsertUpdateFeiras
	feirasRepository feiraInterfaces.FeirasRepository
	feirasUsecases   feiraInterfaces.FeirasUsecases
)

func initFeiraHandler() {
	var dbConn database.DbConnect
	db := dbConn.ConnectDB()

	feirasRepository = feiraRepository.NewFeirasRepository(db)
	feirasUsecases = feiraUsecases.NewFeirasUsecases(feirasRepository)
}

// @Summary Feiras
// @Description Feiras API
// @Param data body models.SearchFeira true "body request"
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 500
// @Router /api/feiras/search [post]
func GetFeiraSearch(c *fiber.Ctx) error {

	initFeiraHandler()

	// Store the body in the note and return error if encountered
	err := c.BodyParser(&SearchFeira)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	if SearchFeira.Pagina == "" {
		SearchFeira.Pagina = "1"
	}

	paging, err := feirasUsecases.GetFeiraSearch(SearchFeira)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// If no note is present return an error
	if paging.TotalPages == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Feiras present", "data": nil})
	}

	// Else return Feiras
	return c.JSON(fiber.Map{"status": "success", "message": "Found data", "data": paging})
}

// @Summary Feiras
// @Description Feiras API
// @Param page path integer true "page number"
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 404
// @Router /api/feiras/p/{page} [get]
func GetFeirasPagination(c *fiber.Ctx) error {

	initFeiraHandler()

	paging, _ := feirasUsecases.GetFeirasPagination(c.Params("page", "1"))

	// If no note is present return an error
	if paging.TotalPages == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Feiras present", "data": nil})
	}

	// Else return Feiras
	return c.JSON(fiber.Map{"status": "success", "message": "Found data", "data": paging})
}

// @Summary Feiras
// @Description Feiras API
// @Param id path string true "id to search"
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 404
// @Router /api/feiras/{id} [get]
func GetFeira(c *fiber.Ctx) error {

	initFeiraHandler()

	id := c.Params("id")

	feira, _ := feirasUsecases.GetFeira(id)
	if feira.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Feiras present", "data": nil})
	}

	// Return the Feira with the Id
	return c.JSON(fiber.Map{"status": "success", "message": "Found data", "data": feira})
}

// @Summary Feiras
// @Description Feiras API
// @Accept  json
// @Produce  json
// @Param data body models.FeiraLivre true "body request"
// @Success 200
// @Failure 500
// @Router /api/feiras [post]
func CreateFeira(c *fiber.Ctx) error {

	initFeiraHandler()

	// Store the body in the note and return error if encountered
	err := c.BodyParser(&Feira)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Create the Feira and return error if encountered
	feira, err := feirasUsecases.CreateFeira(Feira)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Feira - " + fmt.Sprintf("%s", err), "data": nil})
	}

	// Return the created Feira
	return c.JSON(fiber.Map{"status": "success", "message": "Created data", "data": feira})
}

// @Summary Feiras
// @Description Feiras API
// @Param data body models.FeiraLivre true "body request"
// @Param id path string true "id to update"
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 404
// @Router /api/feiras/{id} [put]
func UpdateFeira(c *fiber.Ctx) error {

	initFeiraHandler()

	// Read the param FeiraId
	id := c.Params("id", "")

	// If no such Feira present return an error
	if id == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Feira present", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	err := c.BodyParser(&Feira)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input - " + fmt.Sprintf("%s", err), "data": nil})
	}

	// Save the Changes
	inUpFeira, err := feirasUsecases.UpdateFeira(id, Feira)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not Update Feira - " + fmt.Sprintf("%s", err), "data": nil})
	}

	// Return the updated Feira
	return c.JSON(fiber.Map{"status": "success", "message": "Updated data", "data": inUpFeira})
}

// @Summary Feira
// @Description Feira API
// @Param id path string true "id to delete"
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 404
// @Router /api/feiras/{id} [delete]
func DeleteFeira(c *fiber.Ctx) error {

	initFeiraHandler()

	// Read the param FeiraId
	id := c.Params("id", "")

	// If no such Feira present return an error
	if id == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Feira present", "data": nil})
	}

	// Delete the Feira and return error if encountered
	err := feirasUsecases.DeleteFeira(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete Feira - " + fmt.Sprintf("%s", err), "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Feira"})
}
