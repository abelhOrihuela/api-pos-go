package domain

import (
	"gorm.io/gorm"
	"pos.com/app/db"
	"pos.com/app/errs"
)

type Settings struct {
	gorm.Model
	Id        int    `gorm:"primaryKey;autoIncrement" db:"id"`
	Name      string `gorm:"not null;type:varchar(100);default:null" db:"name"`
	VersionDB string `gorm:"not null;type:varchar(100);default:null" db:"version_db"`
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
