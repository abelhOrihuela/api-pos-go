package dto

import (
	"pos.com/app/errs"
)

type Product struct {
	Id          uint    `json:"id"`
	Uuid        string  `json:"uuid"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Barcode     string  `json:"barcode"`
	Price       float64 `json:"price"`
}

type ProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Barcode     string  `json:"barcode"`
	Price       float64 `json:"price"`
	Category    uint    `json:"category"`
}

func (p ProductRequest) Validate() *errs.AppError {
	if p.Price <= 0 {
		return errs.NewValidationError("El precio de un producto noe puede ser menor o igual a cero")
	}
	if p.Name == "" {
		return errs.NewValidationError("El nombre del producto es requerido")
	}
	return nil
}
