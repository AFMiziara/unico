package feiraRepositories

import (
	"os"
	"testing"

	"github.com/fsvxavier/unico/internal/models"
	"github.com/fsvxavier/unico/pkg/enviroment"
	"github.com/stretchr/testify/assert"
)

var (
	fr FeirasRepository
)

func SetupTests() {

	dir, _ := os.Getwd()
	os.Setenv("ENV", "production")
	var env enviroment.ConfigEnviroment
	env.SetFileConfig(dir + "/../../../config/env.json")
	env.GetTag("ENV")

	os.Setenv("EXECUTE_MIGRATION", "FALSE")

}

func TestGetFeiraSearch(t *testing.T) {

	SetupTests()

	mockSearchData := models.SearchFeira{
		Pagina:    "1",
		Bairro:    "FORMOSA",
		Distrito:  "ARICANDUVA",
		Regiao5:   "Leste",
		NomeFeira: "RECORD",
	}
	t.Run("sucess_GetFeiraSearch", func(t *testing.T) {

		got, err := fr.GetFeiraSearch(mockSearchData)

		assert.NoError(t, err)
		assert.NotNil(t, got.Rows)
	})

	t.Run("error_GetFeiraSearch", func(t *testing.T) {

		mockSearchData.Pagina = ""

		_, err := fr.GetFeiraSearch(mockSearchData)

		assert.Error(t, err)
	})
}

func TestGetFeirasPagination(t *testing.T) {

	SetupTests()

	t.Run("sucess", func(t *testing.T) {

		page := "1"

		got, err := fr.GetFeirasPagination(page)

		assert.NoError(t, err)
		assert.NotNil(t, got)
	})

	t.Run("error", func(t *testing.T) {

		_, err := fr.GetFeirasPagination("")

		assert.Error(t, err)
	})
}

func TestGetFeira(t *testing.T) {

	SetupTests()

	t.Run("sucess", func(t *testing.T) {

		id := "9658c7f9-d65b-406b-9c6f-c19b0c3de8a8"

		got, err := fr.GetFeira(id)

		assert.NoError(t, err)
		assert.NotNil(t, got)
	})

	t.Run("empty", func(t *testing.T) {

		got, err := fr.GetFeira("\\\\\\")

		assert.NoError(t, err)
		assert.NotNil(t, got)
	})
}
