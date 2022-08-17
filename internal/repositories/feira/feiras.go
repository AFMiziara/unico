package feiraRepositories

import (
	"fmt"
	"strconv"

	"github.com/fsvxavier/unico/database"
	feiraInterfaces "github.com/fsvxavier/unico/internal/interfaces/feira"
	"github.com/fsvxavier/unico/internal/models"
	"github.com/fsvxavier/unico/pkg/logger"
)

type feirasRepository struct{}

var (
	Feiras []models.FeiraLivre
	Feira  models.FeiraLivre
)

// NewFeirasRepository exports an interface to FeirasRepository
func NewFeirasRepository() feiraInterfaces.FeirasRepository {
	return &feirasRepository{}
}

func (f *feirasRepository) GetFeiraSearch(feira models.SearchFeira) (database.Pagination, error) {

	logg := new(logger.GenericLogger)
	logg.Module = "api"
	logg.GetLogger()

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

	logg := new(logger.GenericLogger)
	logg.Module = "api"
	logg.GetLogger()

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
	if err != nil {
		logg.LogIt("FATAL", fmt.Sprintf("[FATAL] - Failed search and pagination data. Erro: %s", err.Error()), nil)
	}
	paging.Rows = Feiras

	return paging, err
}

func (f *feirasRepository) GetFeira(id string) ([]models.FeiraLivre, error) {

	logg := new(logger.GenericLogger)
	logg.Module = "api"
	logg.GetLogger()

	var dbConn database.DbConnect
	db := dbConn.ConnectDB()

	err := db.Find(&Feiras, "id = ?", id).Error
	if err != nil {
		logg.LogIt("FATAL", fmt.Sprintf("[FATAL] - Failed search data. Erro: %s", err.Error()), nil)
	}

	return Feiras, err
}

func (f *feirasRepository) CreateFeira(feira models.FeiraLivre) (models.FeiraLivre, error) {

	logg := new(logger.GenericLogger)
	logg.Module = "api"
	logg.GetLogger()

	var dbConn database.DbConnect
	db := dbConn.ConnectDB()

	err := db.Create(&feira).Error
	if err != nil {
		logg.LogIt("FATAL", fmt.Sprintf("[FATAL] - Failed create data. Erro: %s", err.Error()), nil)
	}

	return feira, err
}

func (f *feirasRepository) UpdateFeira(id string, InUpFeira models.InsertUpdateFeiras) (models.InsertUpdateFeiras, error) {
	var dbConn database.DbConnect
	db := dbConn.ConnectDB()

	err := db.Find(&Feira).Where("id = ?", id).Save(&InUpFeira).Error
	return InUpFeira, err
}

func (f *feirasRepository) DeleteFeira(id string) error {
	var dbConn database.DbConnect
	db := dbConn.ConnectDB()

	err := db.Delete(&Feira, "id = ?", id).Error
	return err
}
