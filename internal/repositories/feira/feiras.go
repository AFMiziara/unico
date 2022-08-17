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

	var dbConn database.DbConnect
	db := dbConn.ConnectDB()

	numPage, err := strconv.Atoi(feira.Pagina)
	if err != nil {
		return database.Pagination{}, err
	}

	paging := database.Pagination{
		Page: numPage,
	}

	if feira.Distrito != "" {
		db = db.Where("distrito like ?", "%"+feira.Distrito+"%")
	}

	if feira.Regiao5 != "" {
		db = db.Where("regiao5 like ?", "%"+feira.Regiao5+"%")
	}

	if feira.NomeFeira != "" {
		db = db.Where("nome_feira like ?", "%"+feira.NomeFeira+"%")
	}

	if feira.Bairro != "" {
		db = db.Where("bairro like ?", "%"+feira.Bairro+"%")
	}

	err = db.Scopes(paging.Paginates(Feiras, db)).Find(&Feiras).Error

	paging.Rows = Feiras

	return paging, err
}

func (f *feirasRepository) GetFeirasPagination(page string) (database.Pagination, error) {

	var dbConn database.DbConnect
	db := dbConn.ConnectDB()

	numPage, err := strconv.Atoi(page)
	if err != nil {
		return database.Pagination{}, err
	}

	paging := database.Pagination{
		Page: numPage,
	}

	err = db.Scopes(paging.Paginates(Feiras, db)).Find(&Feiras).Error
	paging.Rows = Feiras

	return paging, err
}

func (f *feirasRepository) GetFeira(id string) ([]models.FeiraLivre, error) {

	var dbConn database.DbConnect
	db := dbConn.ConnectDB()

	err := db.Find(&Feiras, "id = ?", id).Error

	return Feiras, err
}

func (f *feirasRepository) CreateFeira(Feira models.FeiraLivre) (models.FeiraLivre, error) {

	var dbConn database.DbConnect
	db := dbConn.ConnectDB()

	db.Create(&Feira)
	return Feira, nil
}

func (f *feirasRepository) UpdateFeira(id string, InUpFeira models.FeiraLivre) (models.FeiraLivre, error) {
	var dbConn database.DbConnect
	db := dbConn.ConnectDB()

	InUpFeira.ID = id

	db.Find(&Feira).Where("id = ?", id).Save(&InUpFeira)
	return InUpFeira, nil
}

func (f *feirasRepository) DeleteFeira(id string) error {
	var dbConn database.DbConnect
	db := dbConn.ConnectDB()

	err := db.Delete(&Feira, "id = ?", id).Error
	return err
}
