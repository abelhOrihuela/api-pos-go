package domain

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"pos.com/app/db"
	"pos.com/app/dto"
	"pos.com/app/errs"
)

type User struct {
	gorm.Model
	Id       uint   `gorm:"primaryKey;autoIncrement" db:"id"`
	Uuid     string `db:"uuid"`
	Email    string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

func CreateUser(req dto.UserRequest) (*User, *errs.AppError) {

	u := User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	err := db.Database.Create(&u).Error

	if err != nil {
		return nil, errs.NewUnexpectedDatabaseError("Unexpected error during the creation of user" + err.Error())
	}
	return &u, nil

}

func (u User) ToDto() dto.UserResponse {
	return dto.UserResponse{
		Username: u.Username,
		Email:    u.Email,
		Role:     u.Role,
	}
}

func (user *User) BeforeSave(*gorm.DB) error {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
