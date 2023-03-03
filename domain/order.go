package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Id         uint    `gorm:"primaryKey;autoIncrement" db:"id"`
	Uuid       string  `db:"uuid"`
	Amount     float64 `db:"name"`
	TotalItems int     `db:"price"`
}
