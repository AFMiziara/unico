package feiraRepositories

import (
	"fmt"
	"strconv"

	"github.com/fsvxavier/unico/database"
	feiraInterfaces "github.com/fsvxavier/unico/internal/interfaces/feira"
	"github.com/fsvxavier/unico/internal/models"
	"github.com/fsvxavier/unico/pkg/logger"
)

type FeirasRepository struct{}

var (
	Feiras []models.FeiraLivre
	Feira  models.FeiraLivre
	dbConn database.DbConnect
	paging database.Pagination
)

// NewFeirasRepository exports an interface to FeirasRepository
func NewFeirasRepository() feiraInterfaces.FeirasRepository {
	return &FeirasRepository{}
}

func (f *FeirasRepository) GetFeiraSearch(feira models.SearchFeira) (database.Pagination, error) {

	logg := new(logger.GenericLogger)
	logg.Module = "api"
	logg.GetLogger()

	db := dbConn.ConnectDB()

	numPage, err := strconv.Atoi(feira.Pagina)
	if err != nil {
		return paging, err
	} else {
		paging.Page = numPage
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

	db.Scopes(paging.Paginates(Feiras, db)).Find(&Feiras)

	paging.Rows = Feiras

	return paging, err
}

func (f *FeirasRepository) GetFeirasPagination(page string) (database.Pagination, error) {

	logg := new(logger.GenericLogger)
	logg.Module = "api"
	logg.GetLogger()

	numPage, err := strconv.Atoi(page)
	if err != nil {
		return paging, err
	} else {
		paging.Page = numPage
	}

	db := dbConn.ConnectDB()

	err = db.Scopes(paging.Paginates(Feiras, db)).Find(&Feiras).Error
	if err != nil {
		logg.LogIt("FATAL", fmt.Sprintf("[FATAL] - Failed search and pagination data. Erro: %s", err.Error()), nil)
	}
	paging.Rows = Feiras

	return paging, err
}

func (f *FeirasRepository) GetFeira(id string) ([]models.FeiraLivre, error) {

	logg := new(logger.GenericLogger)
	logg.Module = "api"
	logg.GetLogger()

	db := dbConn.ConnectDB()

	err := db.Find(&Feiras, "id = ?", id).Error
	if err != nil {
		logg.LogIt("FATAL", fmt.Sprintf("[FATAL] - Failed search data. Erro: %s", err.Error()), nil)
	}

	return Feiras, err
}

func (f *FeirasRepository) CreateFeira(InUpFeira models.InsertUpdateFeiras) (models.InsertUpdateFeiras, error) {

	logg := new(logger.GenericLogger)
	logg.Module = "api"
	logg.GetLogger()

	db := dbConn.ConnectDB()

	err := db.Create(&InUpFeira).Error
	if err != nil {
		logg.LogIt("FATAL", fmt.Sprintf("[FATAL] - Failed create data. Erro: %s", err.Error()), nil)
	}

	return InUpFeira, err
}

func (f *FeirasRepository) UpdateFeira(id string, InUpFeira models.InsertUpdateFeiras) (models.InsertUpdateFeiras, error) {
	db := dbConn.ConnectDB()

	err := db.Find(&Feira).Where("id = ?", id).Save(&InUpFeira).Error
	return InUpFeira, err
}

func (f *FeirasRepository) DeleteFeira(id string) error {
	db := dbConn.ConnectDB()

	err := db.Delete(&Feira, "id = ?", id).Error
	return err
}
