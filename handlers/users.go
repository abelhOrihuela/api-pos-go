package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"pos.com/app/domain"
	"pos.com/app/dto"
)

/*
* Create user
 */
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	var request dto.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	u, err := domain.CreateUser(request)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		return
	}

	WriteResponse(rw, http.StatusOK, u.ToDto())
}

/*
* Get user
 */
func GetUser(rw http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	uuid := requestVars["user_uuid"]

	u, err := domain.FindUserByUuid(uuid)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		return
	}
	WriteResponse(rw, http.StatusOK, u.ToDto())
}

/*
* Update user
 */
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	uuid := requestVars["user_uuid"]
	var request dto.UserRequestUpdate

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	u, err := domain.UpdateUser(uuid, request)
	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		return
	}

	WriteResponse(rw, http.StatusOK, u.ToDto())
}

/*
* Delete user
 */
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	uuid := requestVars["user_uuid"]

	u, err := domain.DeleteUser(uuid)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		return
	}

	WriteResponse(rw, http.StatusOK, u.ToDto())
}

/*
* Get all users
 */
func GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	response := domain.GetAllUsers(r)
	WriteResponse(rw, http.StatusOK, response)
}
