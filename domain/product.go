package domain

import (
	"net/http"
	"strings"

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
	Metadata         string   `gorm:"not null;type:varchar(100);default:''" db:"metadata"`
	Barcode          string   `gorm:"unique;not null;type:varchar(100);default:null" db:"barcode"`
	Price            float64  `gorm:"not null;type:double precision;default:null" db:"price"`
	CurrentExistence float64  `gorm:"not null;type:double precision;default:null" db:"current_existence"`
	Unit             string   `gorm:"not null;type:string;default:null" db:"unit"`
	CategoryID       int      `gorm:"not null;type:int;default:null" db:"category_id"`
	Category         Category `gorm:"foreignKey:CategoryID;references:Id"`
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

func UpdateProduct(uuid string, req dto.ProductRequest) (*Product, *errs.AppError) {
	var product Product
	err := db.Database.First(&product, "uuid = ?", uuid).Error
	if err != nil {
		return nil, errs.NewDefaultError(err.Error())
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Barcode != "" {
		product.Barcode = req.Barcode
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.CategoryID > 0 {
		product.CategoryID = req.CategoryID
	}
	if req.Unit != "" {
		product.Unit = req.Unit
	}
	if req.CurrentExistence >= 0 {
		product.CurrentExistence = req.CurrentExistence
	}

	if err := db.Database.Save(&product).Error; err != nil {
		return nil, errs.NewDefaultError(err.Error())
	}

	return &product, nil
}

func FindProductByUuid(uuid string) (*Product, *errs.AppError) {
	var product Product
	err := db.Database.Where(&Product{Uuid: uuid}).First(&product).Error

	if err != nil {
		return nil, errs.NewDefaultError(err.Error())
	}
	return &product, nil
}

func DeleteProduct(uuid string) (*Product, *errs.AppError) {
	var product Product

	err := db.Database.Where("uuid = ?", uuid).Delete(&product).Error
	if err != nil {
		return nil, errs.NewDefaultError(err.Error())
	}

	return &product, nil
}

func GetAllProducts(req *http.Request) paginate.Page {
	model := db.Database.Where("deleted_at IS NULL").Preload("Category").Model(&Product{})
	pg := paginate.New()

	page := pg.With(model).Request(req).Response(&[]dto.Product{})

	return page
}

func Search(query string) []Product {
	termSearch := strings.Join(strings.Split(strings.ToLower(query), " "), "-")
	p := make([]Product, 0)
	db.Database.Where("barcode LIKE ? OR name LIKE ? OR metadata LIKE ?", "%"+termSearch+"%", "%"+termSearch+"%", "%"+termSearch+"%").Find(&p)

	return p
}

func (product *Product) BeforeCreate(*gorm.DB) error {

	product.Uuid = uuid.NewString()
	return nil
}

func (product *Product) BeforeSave(*gorm.DB) error {

	name := strings.Join(strings.Split(strings.ToLower(product.Name), " "), "-")
	description := strings.Join(strings.Split(strings.ToLower(product.Description), " "), "-")

	product.Metadata = product.Barcode + "," + name + "," + description
	return nil
}

func (p Product) ToDto() dto.SingleProduct {

	deletedAt := ""

	if p.DeletedAt.Valid {
		deletedAt = p.DeletedAt.Time.String()
	}

	return dto.SingleProduct{
		Uuid:             p.Uuid,
		Id:               p.Id,
		Name:             p.Name,
		Description:      p.Description,
		Unit:             p.Unit,
		Barcode:          p.Barcode,
		Price:            p.Price,
		CategoryID:       p.CategoryID,
		CurrentExistence: p.CurrentExistence,
		CreatedAt:        p.CreatedAt.String(),
		UpdatedAt:        p.UpdatedAt.String(),
		DeletedAt:        deletedAt,
	}
}
