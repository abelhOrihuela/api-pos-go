package handlers

import (
	"net/http"

	"github.com/mrz1836/go-sanitize"
	"pos.com/app/domain"
	"pos.com/app/dto"
)

func GetAllProducts(rw http.ResponseWriter, r *http.Request) {
	var response []dto.ProductResponse
	c := domain.GetAll()

	for _, cs := range c {
		response = append(response, cs.ToDto())
	}

	WriteResponse(rw, http.StatusOK, response)
}

func Search(rw http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	var response []dto.ProductResponse
	c := domain.Search(sanitize.Scripts(q))

	for _, cs := range c {
		response = append(response, cs.ToDto())
	}

	WriteResponse(rw, http.StatusOK, response)
}
