package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

func GetOrder(rw http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	uuid := requestVars["order_uuid"]
	order, err := domain.GetOrder(uuid)

	if err != nil {
		WriteResponse(rw, http.StatusNotFound, err.AsMessage())
		return
	}

	WriteResponse(rw, http.StatusOK, order.ToDto())
}

func GetOrders(rw http.ResponseWriter, r *http.Request) {
	response := domain.GetAllOrders(r)
	WriteResponse(rw, http.StatusOK, response)
}
