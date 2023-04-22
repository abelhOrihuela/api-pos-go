package domain

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
	"pos.com/app/db"
	"pos.com/app/dto"
	"pos.com/app/errs"
)

type Product struct {
	gorm.Model
	Id          uint    `gorm:"primaryKey;autoIncrement" db:"id"`
	Uuid        string  `gorm:"unique;not null;type:varchar(100)" db:"uuid"`
	Name        string  `gorm:"not null;type:varchar(100)" db:"name"`
	Description string  `gorm:"not null;type:varchar(100)" db:"description"`
	Barcode     string  `gorm:"unique;not null;type:varchar(100)" db:"barcode"`
	Price       float64 `gorm:"not null;type:double" db:"price"`
	Category    uint    `gorm:"not null;type:int" db:"category_id"`
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
		Category:    req.Category,
	}
	db.Database.Create(&p)

	return &p, nil

}

func GetAllProducts(req *http.Request) paginate.Page {
	model := db.Database.Where("price IS NOT NULL").Model(&Product{})
	pg := paginate.New()

	page := pg.With(model).Request(req).Response(&[]dto.Product{})

	return page
}

func Search(query string) []Product {
	p := make([]Product, 0)
	db.Database.Where("barcode LIKE ? OR name LIKE ?", "%"+query+"%", "%"+query+"%").Find(&p)

	return p
}

func (product *Product) BeforeSave(*gorm.DB) error {

	product.Uuid = uuid.NewString()
	return nil
}

func (p Product) ToDto() dto.Product {
	return dto.Product{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Barcode:     p.Barcode,
		Price:       p.Price,
	}
}
