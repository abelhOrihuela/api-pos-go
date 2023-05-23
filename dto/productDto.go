package dto

import (
	"pos.com/app/errs"
)

// product response with relationships
type Product struct {
	Id               int      `json:"id"`
	Uuid             string   `json:"uuid"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Barcode          string   `json:"barcode"`
	Price            float64  `json:"price"`
	CategoryID       int      `json:"category_id"`
	Category         Category `json:"category" gorm:"foreignKey:Id;references:CategoryID"`
	Unit             string   `json:"unit"`
	CurrentExistence float64  `json:"current_existence"`
	CreatedAt        string   `json:"created_at"`
	UpdatedAt        string   `json:"updated_at"`
	DeletedAt        string   `json:"deleted_at"`
}

// product response
type SingleProduct struct {
	Id               int     `json:"id"`
	Uuid             string  `json:"uuid"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Barcode          string  `json:"barcode"`
	Unit             string  `json:"unit"`
	Price            float64 `json:"price"`
	CurrentExistence float64 `json:"current_existence"`
	CategoryID       int     `json:"category_id"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	DeletedAt        string  `json:"deleted_at"`
}

// ProductRequest to create a product
type ProductRequest struct {
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Barcode          string  `json:"barcode"`
	Price            float64 `json:"price"`
	CategoryID       int     `json:"category"`
	Unit             string  `json:"unit"`
	CurrentExistence float64 `json:"current_existence"`
}

// Validator of ProductRequest
func (p ProductRequest) Validate() *errs.AppError {
	if p.Barcode == "" {
		return errs.NewValidationError("¡El campo código de barras es requerido!")
	}
	if p.Name == "" {
		return errs.NewValidationError("¡El campo nombre es requerido!")
	}
	if p.Description == "" {
		return errs.NewValidationError("¡El campo descripción es requerido!")
	}
	if p.Unit == "" {
		return errs.NewValidationError("¡El campo unidad es requerido!")
	}
	if p.CategoryID <= 0 {
		return errs.NewValidationError("¡El campo categoria es requerido!")
	}
	if p.Price <= 0 {
		return errs.NewValidationError("El precio de un producto no puede ser menor o igual a cero")
	}

	return nil
}
