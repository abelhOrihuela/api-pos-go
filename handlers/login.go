package handlers

import (
	"encoding/json"
	"net/http"

	"pos.com/app/domain"
	"pos.com/app/dto"
	"pos.com/app/helpers"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	// Parse response
	var request dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.Error())
	} else {

		err := request.ValidateLoginRequest()

		if err != nil {
			WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		} else {
			u, err := domain.FindUserByEmail(request)

			// User not found
			if err != nil {
				WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
			} else {

				// Validate password
				err = u.ValidatePassword(request.Password)

				// Wrong password
				if err != nil {
					WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
				} else {
					// generate token
					token, errToken := helpers.GenerateJWT(u)

					// error token
					if errToken != nil {
						WriteResponse(rw, http.StatusBadRequest, errToken.AsMessage())
					} else {
						// response data with token
						r := dto.TokenResponse{
							AccessToken: token,
						}
						WriteResponse(rw, http.StatusOK, r)
					}
				}
			}
		}

	}
}
