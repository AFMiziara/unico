package feiraInterfaces

import (
	"github.com/fsvxavier/unico/database"
	"github.com/fsvxavier/unico/internal/models"
)

//FeirasRepository ...
type FeirasRepository interface {
	GetFeirasPagination(page string) (database.Pagination, error)
	GetFeira(id string) ([]models.FeiraLivre, error)
	CreateFeira(feira models.FeiraLivre) (models.FeiraLivre, error)
	UpdateFeira(id string, feira models.FeiraLivre) (models.FeiraLivre, error)
	DeleteFeira(id string) error
	GetFeiraSearch(feira models.SearchFeira) (database.Pagination, error)
}

//BancoRepository ...
type FeirasUsecases interface {
	GetFeirasPagination(page string) (database.Pagination, error)
	GetFeira(id string) ([]models.FeiraLivre, error)
	CreateFeira(feira models.FeiraLivre) (models.FeiraLivre, error)
	UpdateFeira(id string, feira models.FeiraLivre) (models.FeiraLivre, error)
	DeleteFeira(id string) error
	GetFeiraSearch(feira models.SearchFeira) (database.Pagination, error)
}
