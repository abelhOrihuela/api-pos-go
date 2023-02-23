package domain

import (
	"gorm.io/gorm"
	"pos.com/app/db"
	"pos.com/app/dto"
)

type Product struct {
	gorm.Model
	Id          uint    `gorm:"primaryKey;autoIncrement" db:"id"`
	Uuid        string  `db:"uuid"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Barcode     string  `db:"barcode"`
	Price       float64 `db:"price"`
}

func Search(query string) []Product {
	p := make([]Product, 0)

	db.Database.Find(&p)
	db.Database.Where("barcode = ?", query).Find(&p)
	return p
}

func GetAll() []Product {
	c := make([]Product, 0)

	db.Database.Find(&c)

	return c
}

func (p Product) ToDto() dto.ProductResponse {
	return dto.ProductResponse{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Barcode:     p.Barcode,
		Price:       p.Price,
	}
}
