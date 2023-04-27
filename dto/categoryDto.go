package dto

import "pos.com/app/errs"

type CategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Category struct {
	Id          int    `json:"id"`
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (c CategoryRequest) Validate() *errs.AppError {
	if c.Name == "" {
		return errs.NewValidationError("El campo nombre es requerido")
	}
	if c.Description == "" {
		return errs.NewValidationError("El campo descripción es requerido")
	}
	return nil
}
