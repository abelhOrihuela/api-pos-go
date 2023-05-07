package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"pos.com/app/domain"
	"pos.com/app/dto"
)

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	var request dto.UserRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.Error())
	} else {
		u, err := domain.CreateUser(request)

		if err != nil {
			WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		} else {
			WriteResponse(rw, http.StatusOK, u.ToDto())
		}
	}
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	userUuid := requestVars["user_uuid"]

	u, err := domain.FindUserByUuid(userUuid)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
	} else {

		WriteResponse(rw, http.StatusOK, u.ToDto())

	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	requestVars := mux.Vars(r)
	userUuid := requestVars["user_uuid"]
	var request dto.UserRequestUpdate
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.Error())
	} else {
		u, err := domain.UpdateUser(request, userUuid)

		if err != nil {
			WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		} else {
			WriteResponse(rw, http.StatusOK, u.ToDto())
		}
	}
}

func GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	response := domain.GetAllUsers(r)
	WriteResponse(rw, http.StatusOK, response)
}
