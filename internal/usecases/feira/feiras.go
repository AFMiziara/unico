package feiraUsecases

import (
	"github.com/fsvxavier/unico/database"
	feiraInterfaces "github.com/fsvxavier/unico/internal/interfaces/feira"
	"github.com/fsvxavier/unico/internal/models"
)

type FeirasUsecases struct {
	FeirasRepository feiraInterfaces.FeirasRepository
}

// NewMySQLFeiraLivreRepository exports an interface to arquivoRepository
func NewFeirasUsecases(i feiraInterfaces.FeirasRepository) feiraInterfaces.FeirasUsecases {
	return &FeirasUsecases{FeirasRepository: i}
}

func (f *FeirasUsecases) GetFeirasPagination(page string) (database.Pagination, error) {
	return f.FeirasRepository.GetFeirasPagination(page)
}

func (t *FeirasUsecases) GetFeiraSearch(feira models.SearchFeira) (database.Pagination, error) {
	return t.FeirasRepository.GetFeiraSearch(feira)
}

func (f *FeirasUsecases) GetFeira(id string) ([]models.FeiraLivre, error) {
	return f.FeirasRepository.GetFeira(id)
}

func (f *FeirasUsecases) CreateFeira(feira models.InsertUpdateFeiras) (models.InsertUpdateFeiras, error) {
	return f.FeirasRepository.CreateFeira(feira)
}

func (f *FeirasUsecases) UpdateFeira(id string, feira models.InsertUpdateFeiras) (models.InsertUpdateFeiras, error) {
	return f.FeirasRepository.UpdateFeira(id, feira)
}

func (f *FeirasUsecases) DeleteFeira(id string) error {
	return f.FeirasRepository.DeleteFeira(id)
}
