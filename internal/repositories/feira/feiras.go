package feiraRepositories

import (
	"fmt"
	"strconv"

	"github.com/fsvxavier/unico/database"
	feiraInterfaces "github.com/fsvxavier/unico/internal/interfaces/feira"
	"github.com/fsvxavier/unico/internal/models"
	"github.com/fsvxavier/unico/pkg/logger"
	"gorm.io/gorm"
)

type FeirasRepository struct {
	db *gorm.DB
}

var Feiras []models.FeiraLivre
var Feira models.FeiraLivre

// NewFeirasRepository exports an interface to FeirasRepository
func NewFeirasRepository(db *gorm.DB) feiraInterfaces.FeirasRepository {
	return &FeirasRepository{db: db}
}

func (f *FeirasRepository) GetFeiraSearch(feira models.SearchFeira) (database.Pagination, error) {

	logg := new(logger.GenericLogger)
	logg.Module = "api"
	logg.GetLogger()

	paging := database.Pagination{}

	numPage, err := strconv.Atoi(feira.Pagina)
	if err != nil {
		return paging, err
	} else {
		paging.Page = numPage
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

func (f *FeirasRepository) GetFeirasPagination(page string) (database.Pagination, error) {

	logg := new(logger.GenericLogger)
	logg.Module = "api"
	logg.GetLogger()

	paging := database.Pagination{}

	numPage, err := strconv.Atoi(page)
	if err != nil {
		return paging, err
	} else {
		paging.Page = numPage
	}

	var dbConn database.DbConnect
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

	var dbConn database.DbConnect
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

	var dbConn database.DbConnect
	db := dbConn.ConnectDB()

	err := db.Create(&InUpFeira).Error
	if err != nil {
		logg.LogIt("FATAL", fmt.Sprintf("[FATAL] - Failed create data. Erro: %s", err.Error()), nil)
	}

	return InUpFeira, err
}

func (f *FeirasRepository) UpdateFeira(id string, InUpFeira models.InsertUpdateFeiras) (models.InsertUpdateFeiras, error) {
	var dbConn database.DbConnect
	db := dbConn.ConnectDB()

	err := db.Find(&Feira).Where("id = ?", id).Save(&InUpFeira).Error
	return InUpFeira, err
}

func (f *FeirasRepository) DeleteFeira(id string) error {
	var dbConn database.DbConnect
	db := dbConn.ConnectDB()

	err := db.Delete(&Feira, "id = ?", id).Error
	return err
}
