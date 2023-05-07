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

type Category struct {
	gorm.Model
	Id          int    `gorm:"primaryKey;autoIncrement" db:"id"`
	Uuid        string `gorm:"unique;not null;type:varchar(100);default:null" db:"uuid"`
	Name        string `gorm:"not null;type:varchar(50);default:null" db:"name"`
	Description string `gorm:"not null;type:varchar(100);default:null" db:"description"`
}

func CreateCategory(req dto.CategoryRequest) (*Category, *errs.AppError) {

	errValidation := req.Validate()

	if errValidation != nil {
		return nil, errValidation
	}

	c := Category{
		Name:        req.Name,
		Description: req.Description,
	}

	err := db.Database.Create(&c).Error

	if err != nil {
		return nil, errs.NewUnexpectedDatabaseError("Unexpected error during the creation of product" + err.Error())
	}
	return &c, nil

}

func GetAllCategories(req *http.Request) paginate.Page {
	model := db.Database.Where("deleted_at IS NULL").Model(&Category{})
	pg := paginate.New()

	page := pg.With(model).Request(req).Response(&[]dto.Category{})

	return page
}

func (category *Category) BeforeCreate(*gorm.DB) error {
	category.Uuid = uuid.NewString()
	return nil
}

func (c Category) ToDto() dto.Category {
	return dto.Category{
		Id:          c.Id,
		Uuid:        c.Uuid,
		Name:        c.Name,
		Description: c.Description,
	}
}
