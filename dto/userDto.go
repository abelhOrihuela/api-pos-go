package dto

import "pos.com/app/errs"

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserRequestUpdate struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserResponse struct {
	Uuid      string `json:"uuid"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type LoginRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (l LoginRequest) ValidateLoginRequest() *errs.AppError {
	if l.Email == "" {
		return errs.NewValidationError("¡Email icnocrrecto!")
	}
	if l.Password == "" {
		return errs.NewValidationError("¡Contraseña requerida!")
	}
	return nil
}
