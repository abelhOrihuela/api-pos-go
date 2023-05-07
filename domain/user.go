package domain

import (
	"html"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/morkid/paginate"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"pos.com/app/db"
	"pos.com/app/dto"
	"pos.com/app/errs"
)

type User struct {
	gorm.Model
	Id       int    `gorm:"primaryKey;autoIncrement" db:"id"`
	Uuid     string `gorm:"unique;not null;type:varchar(100);default:null" db:"uuid"`
	Email    string `gorm:"unique;not null;type:varchar(50);default:null" db:"email"`
	Username string `gorm:"unique;not null;type:varchar(50);default:null" db:"username"`
	Password string `gorm:"not null;type:varchar(100);default:null" db:"password"`
	Role     string `gorm:"not null;type:varchar(20);default:null" db:"role"`
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

func UpdateUser(req dto.UserRequestUpdate, uuid string) (*User, *errs.AppError) {
	var user User

	err := db.Database.Where("uuid =?", uuid).First(&user).Error

	if err != nil {
		return nil, errs.NewDefaultError(err.Error())
	}

	if req.Username != "" {
		user.Username = req.Username
	}

	if req.Role != "" {
		user.Role = req.Role
	}

	db.Database.Save(&user)

	return &user, nil

}

func FindUserByEmail(email string) (*User, *errs.AppError) {
	var user User

	err := db.Database.Where(&User{Email: email}).First(&user).Error

	if err != nil {
		return nil, errs.NewNotFoundError("User not found")
	}
	return &user, nil
}

func FindUserById(id int) (*User, *errs.AppError) {
	var user User

	err := db.Database.Where(&User{Id: id}).First(&user).Error

	if err != nil {
		return nil, errs.NewNotFoundError("User not found")
	}
	return &user, nil
}

func FindUserByUuid(uuid string) (*User, *errs.AppError) {
	var user User

	err := db.Database.Where(&User{Uuid: uuid}).First(&user).Error

	if err != nil {
		return nil, errs.NewNotFoundError("User not found")
	}
	return &user, nil
}

func GetAllUsers(req *http.Request) paginate.Page {
	model := db.Database.Where("deleted_at IS NULL").Model(&User{})
	pg := paginate.New()

	page := pg.With(model).Request(req).Response(&[]dto.UserResponse{})

	return page
}

func (user *User) ValidatePassword(password string) *errs.AppError {

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return errs.NewValidationError("¡Password incorrecto.!")
	}
	return nil
}

func (user *User) BeforeCreate(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.Uuid = uuid.NewString()
	return nil
}

func (u User) ToDto() dto.UserResponse {
	return dto.UserResponse{
		Uuid:     u.Uuid,
		Username: u.Username,
		Email:    u.Email,
		Role:     u.Role,
	}
}
