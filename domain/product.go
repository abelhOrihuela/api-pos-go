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
	Id               int      `gorm:"primaryKey;autoIncrement" db:"id"`
	Uuid             string   `gorm:"unique;not null;type:varchar(100);default:null" db:"uuid"`
	Name             string   `gorm:"not null;type:varchar(100);default:null" db:"name"`
	Description      string   `gorm:"not null;type:varchar(100);default:null" db:"description"`
	Barcode          string   `gorm:"unique;not null;type:varchar(100);default:null" db:"barcode"`
	Price            float64  `gorm:"not null;type:double precision;default:null" db:"price"`
	CurrentExistence int64    `gorm:"not null;type:int;default:0" db:"current_existence"`
	Unit             string   `gorm:"not null;type:string;default:null" db:"unit"`
	CategoryID       int      `gorm:"not null;type:int;default:null" db:"category_id"`
	Category         Category `gorm:"foreignKey:Id;references:CategoryID"`
}

func CreateProduct(req dto.ProductRequest) (*Product, *errs.AppError) {

	errValidation := req.Validate()

	if errValidation != nil {
		return nil, errValidation
	}

	p := Product{
		Barcode:          req.Barcode,
		Name:             req.Name,
		Description:      req.Description,
		Price:            req.Price,
		CategoryID:       req.CategoryID,
		Unit:             req.Unit,
		CurrentExistence: req.CurrentExistence,
	}

	err := db.Database.Create(&p).Error

	if err != nil {
		return nil, errs.NewUnexpectedDatabaseError("Unexpected error during the creation of product" + err.Error())
	}
	return &p, nil

}

func GetAllProducts(req *http.Request) paginate.Page {
	model := db.Database.Where("deleted_at IS NULL").Preload("Category").Model(&Product{})
	pg := paginate.New()

	page := pg.With(model).Request(req).Response(&[]dto.Product{})

	return page
}

func Search(query string) []Product {
	p := make([]Product, 0)
	db.Database.Where("barcode LIKE ? OR name LIKE ?", "%"+query+"%", "%"+query+"%").Find(&p)

	return p
}

func (product *Product) BeforeCreate(*gorm.DB) error {

	product.Uuid = uuid.NewString()
	return nil
}

func (p Product) ToDto() dto.SingleProduct {
	return dto.SingleProduct{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Barcode:     p.Barcode,
		Price:       p.Price,
		CategoryID:  p.CategoryID,
	}
}
