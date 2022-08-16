package feiraRepositories

import (
	"os"
	"testing"

	"github.com/fsvxavier/unico/internal/mocks"
	"github.com/fsvxavier/unico/internal/models"
	"github.com/fsvxavier/unico/pkg/enviroment"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type v2Suite struct {
	db *gorm.DB
}

func setupTest() {

	dir, _ := os.Getwd()
	os.Setenv("ENV", "production")
	var env enviroment.ConfigEnviroment
	env.SetFileConfig(dir + "/../../../config/env.json")
	env.GetTag("ENV")

	os.Setenv("EXECUTE_MIGRATION", "FALSE")

}

func TestGetFeiraSearch(t *testing.T) {

	setupTest()

	db, mock := mocks.NewDatabase()

	mockSearchData := models.SearchFeira{
		Pagina:    "1",
		Bairro:    "FORMOSA",
		Distrito:  "ARICANDUVA",
		Regiao5:   "Leste",
		NomeFeira: "RECORD",
	}

	t.Run("sucess", func(t *testing.T) {
		mock.MatchExpectationsInOrder(false)

		query := (`SELECT count(1) FROM "feira_livre" WHERE distrito like $1 AND regiao5 like $2 AND nome_feira like $3 AND bairro like $4`)

		mock.ExpectQuery(query).WithArgs().
			//WithArgs("%"+mockSearchData.Distrito+"%", "%"+mockSearchData.Regiao5+"%", "%"+mockSearchData.NomeFeira+"%", "%"+mockSearchData.Bairro+"%").
			WillReturnRows()

		a := NewFeirasRepository(db)
		got, err := a.GetFeiraSearch(mockSearchData)

		assert.NoError(t, err)
		assert.NotNil(t, got)
	})

	t.Run("error", func(t *testing.T) {

		mockSearchData.Pagina = ""

		a := NewFeirasRepository(db)
		_, err := a.GetFeiraSearch(mockSearchData)

		assert.Error(t, err)
	})
}

func TestGetFeirasPagination(t *testing.T) {

	setupTest()

	s := &v2Suite{}
	a := &FeirasRepository{s.db}

	t.Run("sucess", func(t *testing.T) {

		page := "1"

		got, err := a.GetFeirasPagination(page)

		assert.NoError(t, err)
		assert.NotNil(t, got)
	})

	t.Run("error", func(t *testing.T) {

		_, err := a.GetFeirasPagination("")

		assert.Error(t, err)
	})
}

func TestGetFeira(t *testing.T) {

	setupTest()

	s := &v2Suite{}
	a := &FeirasRepository{s.db}

	t.Run("sucess", func(t *testing.T) {

		id := "9658c7f9-d65b-406b-9c6f-c19b0c3de8a8"

		got, err := a.GetFeira(id)

		assert.NoError(t, err)
		assert.NotNil(t, got)
	})

	t.Run("empty", func(t *testing.T) {

		got, err := a.GetFeira("\\\\\\")

		assert.NoError(t, err)
		assert.NotNil(t, got)
	})
}
