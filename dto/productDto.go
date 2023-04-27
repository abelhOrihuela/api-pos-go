package dto

import (
	"pos.com/app/errs"
)

// product response with relationships
type Product struct {
	Id          int      `json:"id"`
	Uuid        string   `json:"uuid"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Barcode     string   `json:"barcode"`
	Price       float64  `json:"price"`
	CategoryID  int      `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:Id;references:CategoryID"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

// product response
type SingleProduct struct {
	Id          int     `json:"id"`
	Uuid        string  `json:"uuid"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Barcode     string  `json:"barcode"`
	Price       float64 `json:"price"`
	CategoryID  int     `json:"category_id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// ProductRequest to create a product
type ProductRequest struct {
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Barcode          string  `json:"barcode"`
	Price            float64 `json:"price"`
	CategoryID       int     `json:"category"`
	Unit             string  `json:"unit"`
	CurrentExistence int64   `json:"current_existence"`
}

// Validator of ProductRequest
func (p ProductRequest) Validate() *errs.AppError {
	if p.Price <= 0 {
		return errs.NewValidationError("El precio de un producto noe puede ser menor o igual a cero")
	}
	if p.Name == "" {
		return errs.NewValidationError("El nombre del producto es requerido")
	}
	return nil
}
