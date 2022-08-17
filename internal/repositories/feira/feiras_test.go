package feiraRepositories

import (
	"errors"
	"os"
	"testing"

	"github.com/fsvxavier/unico/database"
	feiraInterfaces "github.com/fsvxavier/unico/internal/interfaces/feira"
	"github.com/fsvxavier/unico/internal/models"
	"github.com/fsvxavier/unico/pkg/enviroment"
	"github.com/stretchr/testify/assert"
)

var (
	fr     feiraInterfaces.FeirasRepository
	idTest string
)

func SetupTests() {

	dir, _ := os.Getwd()
	os.Setenv("ENV", "production")
	var env enviroment.ConfigEnviroment
	env.SetFileConfig(dir + "/../../../config/env.json")
	env.GetTag("ENV")

	os.Setenv("EXECUTE_MIGRATION", "FALSE")

	var dbConn database.DbConnect
	db := dbConn.ConnectDB()

	fr = NewFeirasRepository(db)

}

func TestCreateFeira(t *testing.T) {

	SetupTests()

	mockFeiras := models.FeiraLivre{
		Longi:      -1,
		Lat:        -1,
		Setcens:    1,
		Areap:      1,
		Coddist:    1,
		Codsubpref: 1,
		Distrito:   "asdasd",
		Subprefe:   "asdasd",
		Regiao5:    "asdasd",
		Regiao8:    "asdasd",
		NomeFeira:  "asdasd",
		Registro:   "asdasd",
		Logradouro: "asdasd",
		Numero:     "asdasd",
		Bairro:     "asdasd",
		Referencia: "asdasd",
	}

	t.Run("sucess", func(t *testing.T) {

		var err error = nil

		p, _ := fr.CreateFeira(mockFeiras)
		if p.ID == "" {
			err = errors.New("LastInsertId is empty")
		}

		idTest = p.ID

		assert.NoError(t, err)
		assert.NotNil(t, p)
	})
}

func TestUpdateFeira(t *testing.T) {

	SetupTests()

	mockFeiras := models.FeiraLivre{
		Longi:      0,
		Lat:        0,
		Setcens:    0,
		Areap:      0,
		Coddist:    0,
		Distrito:   "",
		Codsubpref: 0,
		Subprefe:   "",
		Regiao5:    "",
		Regiao8:    "",
		NomeFeira:  "",
		Registro:   "",
		Logradouro: "",
		Numero:     "",
		Bairro:     "",
		Referencia: "",
	}

	t.Run("sucess", func(t *testing.T) {

		var err error = nil

		c, _ := fr.CreateFeira(mockFeiras)
		if c.ID == "" {
			err = errors.New("LastInsertId is empty")
		}

		p, _ := fr.UpdateFeira(c.ID, mockFeiras)
		if p.ID == "" {
			err = errors.New("Error Update data")
		}

		assert.NoError(t, err)
		assert.NotNil(t, p)
	})
}

func TestDeleteFeira(t *testing.T) {

	SetupTests()

	mockFeiras := models.FeiraLivre{
		Longi:      0,
		Lat:        0,
		Setcens:    0,
		Areap:      0,
		Coddist:    0,
		Distrito:   "",
		Codsubpref: 0,
		Subprefe:   "",
		Regiao5:    "",
		Regiao8:    "",
		NomeFeira:  "",
		Registro:   "",
		Logradouro: "",
		Numero:     "",
		Bairro:     "",
		Referencia: "",
	}

	t.Run("sucess", func(t *testing.T) {

		c, _ := fr.CreateFeira(mockFeiras)

		err := fr.DeleteFeira(c.ID)

		assert.NoError(t, err)
	})
}

func TestGetFeiraSearch(t *testing.T) {

	SetupTests()

	mockSearchData := models.SearchFeira{
		Pagina:    "1",
		Bairro:    "FORMOSA",
		Distrito:  "VILA FORMOSA",
		Regiao5:   "Leste",
		NomeFeira: "MANOEL",
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
