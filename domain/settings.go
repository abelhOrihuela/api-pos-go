package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"pos.com/app/db"
	"pos.com/app/errs"
)

type Settings struct {
	gorm.Model
	Id        int    `gorm:"primaryKey;autoIncrement" db:"id"`
	Uuid      string `gorm:"unique;not null;type:varchar(100);default:null" db:"uuid"`
	VersionDB string `gorm:"not null;type:varchar(100);default:'1.0'" db:"version_db"`
	RootUser  string `gorm:"not null;type:varchar(100);default:''" db:"version_db"`
}

func UpdateVersion(version string) (*Settings, *errs.AppError) {
	var settings Settings
	err := db.Database.First(&settings).Error
	if err != nil {
		return nil, errs.NewDefaultError(err.Error())
	}
	settings.VersionDB = version
	db.Database.Save(&settings)
	return &settings, nil
}

func (settings *Settings) BeforeCreate(*gorm.DB) error {

	settings.Uuid = uuid.NewString()
	return nil
}
