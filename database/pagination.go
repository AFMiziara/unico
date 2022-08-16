package database

import (
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int
	Page       int
	Sort       string
	TotalRows  int64
	TotalPages int
	Rows       interface{}
	HasPrev    bool
	HasNext    bool
	PrevPage   int
	NextPage   int
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page < 1 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id asc"
	}
	return p.Sort
}

func (p *Pagination) Paginates(value interface{}, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(&value).Select("count(1)").Count(&totalRows)

	p.TotalRows = totalRows
	p.Calculate()

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSort())
	}
}

// pages start at 1 - not 0
func (p *Pagination) Calculate() {

	// calculate number of pages
	d := float64(p.TotalRows) / float64(p.GetLimit())
	p.TotalPages = int(math.Ceil(d))

	// HasPrev, HasNext?
	p.HasPrev = p.GetPage() > 1
	p.HasNext = p.GetPage() < p.TotalPages

	// calculate them
	if p.HasPrev {
		p.PrevPage = p.GetPage() - 1
		if p.PrevPage == 0 {
			p.PrevPage = 1
		}
	}
	if p.HasNext {
		p.NextPage = p.Page + 1
		if p.NextPage > p.TotalPages {
			p.NextPage = p.TotalPages
		}
	}
}
