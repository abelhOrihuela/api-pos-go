package handlers

import (
	"encoding/json"
	"net/http"

	"pos.com/app/domain"
	"pos.com/app/dto"
	"pos.com/app/helpers"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	var request dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.ValidateLoginRequest(); err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		return
	}

	u, err := domain.FindUserByEmail(request.Email)
	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		return
	}

	if err := u.ValidatePassword(request.Password); err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		return
	}

	token, err := helpers.GenerateJWT(u)
	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		return
	}

	res := dto.TokenResponse{AccessToken: token}

	WriteResponse(rw, http.StatusOK, res)
}

func Me(rw http.ResponseWriter, r *http.Request) {
	u, err := helpers.CurrentUser(rw, r)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
	} else {
		WriteResponse(rw, http.StatusOK, u.ToDto())
	}
}
