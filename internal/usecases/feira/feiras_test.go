package feiraUsecases

import (
	"testing"

	"github.com/fsvxavier/unico/database"
	mocksfeira "github.com/fsvxavier/unico/internal/mocks/interfaces/feira"
	"github.com/fsvxavier/unico/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGetFeira(t *testing.T) {
	mockFeiras := []models.FeiraLivre{
		{
			ID:         "512b1e41-fe0f-4012-832e-a88693013483",
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
		},
	}
	mockFeirasRepository := new(mocksfeira.FeirasRepository)

	t.Run("sucess", func(t *testing.T) {

		id := "512b1e41-fe0f-4012-832e-a88693013483"
		mockFeirasRepository.On("GetFeira", id).Return(mockFeiras, nil)
		u := NewFeirasUsecases(mockFeirasRepository)
		p, err := u.GetFeira(id)

		assert.NoError(t, err)
		assert.NotNil(t, p)
		mockFeirasRepository.AssertExpectations(t)
	})
}

func TestGetFeiraSearch(t *testing.T) {
	mockFeiras := database.Pagination{
		Limit:      1,
		Page:       1,
		Sort:       "ASC",
		TotalRows:  1,
		TotalPages: 1,
		Rows: []models.FeiraLivre{
			{
				ID:         "512b1e41-fe0f-4012-832e-a88693013483",
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
			},
		},
		HasPrev:  false,
		HasNext:  false,
		PrevPage: 1,
		NextPage: 1,
	}
	mockFeirasRepository := new(mocksfeira.FeirasRepository)

	t.Run("sucess", func(t *testing.T) {

		page := "1"
		mockFeirasRepository.On("GetFeirasPagination", page).Return(mockFeiras, nil)
		u := NewFeirasUsecases(mockFeirasRepository)
		p, err := u.GetFeirasPagination(page)

		assert.NoError(t, err)
		assert.NotNil(t, p)
		mockFeirasRepository.AssertExpectations(t)
	})
}

func TestGetFeirasPagination(t *testing.T) {

	mockSearchFeira := models.SearchFeira{
		Pagina:    "1",
		Distrito:  "asdfgçlkjh",
		Regiao5:   "",
		NomeFeira: "",
		Bairro:    "",
	}

	mockFeiras := database.Pagination{
		Limit:      1,
		Page:       1,
		Sort:       "ASC",
		TotalRows:  1,
		TotalPages: 1,
		Rows: []models.FeiraLivre{
			{
				ID:         "512b1e41-fe0f-4012-832e-a88693013483",
				Longi:      0,
				Lat:        0,
				Setcens:    0,
				Areap:      0,
				Coddist:    0,
				Distrito:   "asdfgçlkjh",
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
			},
		},
		HasPrev:  false,
		HasNext:  false,
		PrevPage: 0,
		NextPage: 0,
	}
	mockFeirasRepository := new(mocksfeira.FeirasRepository)

	t.Run("sucess", func(t *testing.T) {

		mockFeirasRepository.On("GetFeiraSearch", mockSearchFeira).Return(mockFeiras, nil)
		u := NewFeirasUsecases(mockFeirasRepository)
		p, err := u.GetFeiraSearch(mockSearchFeira)

		assert.NoError(t, err)
		assert.NotNil(t, p)
		mockFeirasRepository.AssertExpectations(t)
	})
}

func TestCreateFeira(t *testing.T) {
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
	mockFeirasRepository := new(mocksfeira.FeirasRepository)

	t.Run("sucess", func(t *testing.T) {

		mockFeirasRepository.On("CreateFeira", mockFeiras).Return(mockFeiras, nil)
		u := NewFeirasUsecases(mockFeirasRepository)
		p, err := u.CreateFeira(mockFeiras)

		assert.NoError(t, err)
		assert.NotNil(t, p)
		mockFeirasRepository.AssertExpectations(t)
	})
}

func TestUpdateFeira(t *testing.T) {
	mockFeiras := models.InsertUpdateFeiras{
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
	mockFeirasRepository := new(mocksfeira.FeirasRepository)

	t.Run("sucess", func(t *testing.T) {

		id := "512b1e41-fe0f-4012-832e-a88693013483"
		mockFeirasRepository.On("UpdateFeira", id, mockFeiras).Return(mockFeiras, nil)
		u := NewFeirasUsecases(mockFeirasRepository)
		p, err := u.UpdateFeira(id, mockFeiras)

		assert.NoError(t, err)
		assert.NotNil(t, p)
		mockFeirasRepository.AssertExpectations(t)
	})
}

func TestDeleteFeira(t *testing.T) {

	mockFeirasRepository := new(mocksfeira.FeirasRepository)

	t.Run("sucess", func(t *testing.T) {

		id := "512b1e41-fe0f-4012-832e-a88693013483"
		mockFeirasRepository.On("DeleteFeira", id).Return(nil)
		u := NewFeirasUsecases(mockFeirasRepository)
		err := u.DeleteFeira(id)

		assert.NoError(t, err)
		mockFeirasRepository.AssertExpectations(t)
	})
}
