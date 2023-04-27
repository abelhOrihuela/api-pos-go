package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mrz1836/go-sanitize"
	"pos.com/app/domain"
	"pos.com/app/dto"
)

func CreateProduct(rw http.ResponseWriter, r *http.Request) {
	var request dto.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.Error())
	} else {
		p, err := domain.CreateProduct(request)

		if err != nil {
			WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		} else {
			WriteResponse(rw, http.StatusOK, p.ToDto())
		}
	}

}

func GetProducts(rw http.ResponseWriter, r *http.Request) {
	response := domain.GetAllProducts(r)
	WriteResponse(rw, http.StatusOK, response)
}

func Search(rw http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	var response []dto.SingleProduct
	c := domain.Search(sanitize.Scripts(q))

	for _, cs := range c {
		response = append(response, cs.ToDto())
	}

	WriteResponse(rw, http.StatusOK, response)
}
