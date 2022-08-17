package feiraRepositories

import (
	"strconv"

	"github.com/fsvxavier/unico/database"
	feiraInterfaces "github.com/fsvxavier/unico/internal/interfaces/feira"
	"github.com/fsvxavier/unico/internal/models"
	"gorm.io/gorm"
)

type feirasRepository struct {
	db *gorm.DB
}

var (
	Feiras []models.FeiraLivre
	Feira  models.FeiraLivre
)

// NewFeirasRepository exports an interface to FeirasRepository
func NewFeirasRepository(db *gorm.DB) feiraInterfaces.FeirasRepository {
	return &feirasRepository{db: db}
}

func (f *feirasRepository) GetFeiraSearch(feira models.SearchFeira) (database.Pagination, error) {

	numPage, err := strconv.Atoi(feira.Pagina)
	if err != nil {
		return database.Pagination{}, err
	}

	paging := database.Pagination{
		Page: numPage,
	}

	if feira.Distrito != "" {
		f.db = f.db.Where("distrito like ?", "%"+feira.Distrito+"%")
	}

	if feira.Regiao5 != "" {
		f.db = f.db.Where("regiao5 like ?", "%"+feira.Regiao5+"%")
	}

	if feira.NomeFeira != "" {
		f.db = f.db.Where("nome_feira like ?", "%"+feira.NomeFeira+"%")
	}

	if feira.Bairro != "" {
		f.db = f.db.Where("bairro like ?", "%"+feira.Bairro+"%")
	}

	err = f.db.Scopes(paging.Paginates(Feiras, f.db)).Find(&Feiras).Error

	paging.Rows = Feiras

	return paging, err
}

func (f *feirasRepository) GetFeirasPagination(page string) (database.Pagination, error) {

	numPage, err := strconv.Atoi(page)
	if err != nil {
		return database.Pagination{}, err
	}

	paging := database.Pagination{
		Page: numPage,
	}

	err = f.db.Scopes(paging.Paginates(Feiras, f.db)).Find(&Feiras).Error
	paging.Rows = Feiras

	return paging, err
}

func (f *feirasRepository) GetFeira(id string) (models.FeiraLivre, error) {

	err := f.db.Find(&Feira, "id = ?", id).Error

	return Feira, err
}

func (f *feirasRepository) CreateFeira(feira models.FeiraLivre) (models.FeiraLivre, error) {

	feira.ID = ""

	f.db.Create(&feira)
	return Feira, nil
}

func (f *feirasRepository) UpdateFeira(id string, inUpFeira models.FeiraLivre) (models.FeiraLivre, error) {

	inUpFeira.ID = id

	f.db.Model(&Feira).Where("id = ?", id).Updates(inUpFeira)
	return inUpFeira, nil
}

func (f *feirasRepository) DeleteFeira(id string) error {

	err := f.db.Delete(&Feira, "id = ?", id).Error
	return err
}
