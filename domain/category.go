package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Id          uint   `gorm:"primaryKey;autoIncrement" db:"id"`
	Uuid        string `gorm:"unique;not null;type:varchar(100);default:null" db:"uuid"`
	Name        string `gorm:"not null;type:varchar(50);default:null" db:"name"`
	Description string `gorm:"not null;type:varchar(100);default:null" db:"description"`
}

func (category *Category) BeforeSave(*gorm.DB) error {

	category.Uuid = uuid.NewString()
	return nil
}
