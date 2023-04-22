package handlers

import (
	"encoding/json"
	"net/http"

	"pos.com/app/domain"
	"pos.com/app/dto"
)

func CreateOrder(rw http.ResponseWriter, r *http.Request) {
	var request dto.OrderRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.Error())
	} else {
		p, err := domain.CreateOrder(request)

		if err != nil {
			WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		} else {
			WriteResponse(rw, http.StatusOK, p.ToDto())
		}
	}

}

func GetOrders(rw http.ResponseWriter, r *http.Request) {
	response := domain.GetAllOrders(r)
	WriteResponse(rw, http.StatusOK, response)
}
