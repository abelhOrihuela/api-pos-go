package domain

import (
	"gorm.io/gorm"
	"pos.com/app/dto"
)

type Product struct {
	gorm.Model
	Id          uint   `gorm:"primaryKey;autoIncrement" db:"id"`
	Uuid        string `db:"uuid"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Barcode     string `db:"barcode"`
}

func (p Product) ToDto() dto.ProductResponse {
	return dto.ProductResponse{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Barcode:     p.Barcode,
	}
}
