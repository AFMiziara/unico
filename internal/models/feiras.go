package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FeiraLivre struct {
	ID         string `gorm:"default:uuid;primarykey;index" json:"id"`
	Longi      int64  `json:"longi"`
	Lat        int64  `json:"lat"`
	Setcens    int64  `json:"setcens"`
	Areap      int64  `json:"areap"`
	Coddist    int64  `json:"coddist"`
	Distrito   string `json:"distrito"`
	Codsubpref int64  `json:"codsubpref"`
	Subprefe   string `json:"subprefe"`
	Regiao5    string `json:"regiao5"`
	Regiao8    string `json:"regiao8"`
	NomeFeira  string `json:"nome_feira"`
	Registro   string `json:"registro"`
	Logradouro string `json:"logradouro"`
	Numero     string `json:"numero"`
	Bairro     string `json:"bairro"`
	Referencia string `json:"referencia"`
}

func (base *FeiraLivre) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	base.ID = uuid.New().String()
	return
}

type SearchFeira struct {
	Pagina    string
	Distrito  string
	Regiao5   string
	NomeFeira string
	Bairro    string
}

type InsertUpdateFeiras struct {
	Longi      int64  `json:"longi"`
	Lat        int64  `json:"lat"`
	Setcens    int64  `json:"setcens"`
	Areap      int64  `json:"areap"`
	Coddist    int64  `json:"coddist"`
	Distrito   string `json:"distrito"`
	Codsubpref int64  `json:"codsubpref"`
	Subprefe   string `json:"subprefe"`
	Regiao5    string `json:"regiao5"`
	Regiao8    string `json:"regiao8"`
	NomeFeira  string `json:"nome_feira"`
	Registro   string `json:"registro"`
	Logradouro string `json:"logradouro"`
	Numero     string `json:"numero"`
	Bairro     string `json:"bairro"`
	Referencia string `json:"referencia"`
}
