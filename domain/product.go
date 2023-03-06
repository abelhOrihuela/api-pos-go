package domain

import (
	"gorm.io/gorm"
	"pos.com/app/db"
	"pos.com/app/dto"
	"pos.com/app/errs"
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

func CreateProduct(req dto.ProductRequest) (*Product, *errs.AppError) {

	err := req.Validate()

	if err != nil {
		return nil, err
	}

	p := Product{
		Barcode:     req.Barcode,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}
	db.Database.Create(&p)

	return &p, nil

}

func GetAll() []Product {
	c := make([]Product, 0)
	db.Database.Find(&c)

	return c
}

func Search(query string) []Product {
	p := make([]Product, 0)
	db.Database.Where("barcode LIKE ? OR name LIKE ?", "%"+query+"%", "%"+query+"%").Find(&p)

	return p
}

func (p Product) ToDto() dto.ProductResponse {
	return dto.ProductResponse{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Barcode:     p.Barcode,
		Price:       p.Price,
	}
}
