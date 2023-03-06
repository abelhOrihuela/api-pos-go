package handlers

import (
	"encoding/json"
	"net/http"

	"pos.com/app/domain"
	"pos.com/app/dto"
)

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	var request dto.UserRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.Error())
	} else {
		p, err := domain.CreateUser(request)

		if err != nil {
			WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		} else {
			WriteResponse(rw, http.StatusOK, p.ToDto())
		}
	}

}
