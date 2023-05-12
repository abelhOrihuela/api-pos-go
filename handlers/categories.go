package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"pos.com/app/domain"
	"pos.com/app/dto"
)

func CreateCategory(rw http.ResponseWriter, r *http.Request) {
	var request dto.CategoryRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.Error())
	} else {
		c, err := domain.CreateCategory(request)

		if err != nil {
			WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		} else {
			WriteResponse(rw, http.StatusOK, c.ToDto())
		}
	}
}

func UpdateCategory(rw http.ResponseWriter, r *http.Request) {
	var request dto.CategoryRequest

	requestVars := mux.Vars(r)
	uuid := requestVars["category_uuid"]

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.Error())
	} else {
		c, err := domain.UpdateCategory(request, uuid)

		if err != nil {
			WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		} else {
			WriteResponse(rw, http.StatusOK, c.ToDto())
		}
	}
}

func GetAllCategories(rw http.ResponseWriter, r *http.Request) {
	response := domain.GetAllCategories(r)
	WriteResponse(rw, http.StatusOK, response)
}
