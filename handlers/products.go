package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrz1836/go-sanitize"
	"pos.com/app/domain"
	"pos.com/app/dto"
)

/*
* Create product
 */
func CreateProduct(rw http.ResponseWriter, r *http.Request) {
	var request dto.ProductRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	p, err := domain.CreateProduct(request)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		return
	}

	WriteResponse(rw, http.StatusOK, p.ToDto())
}

/*
* Get product
 */
func GetProduct(rw http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	uuid := requestVars["product_uuid"]

	p, err := domain.FindProductByUuid(uuid)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		return
	}

	WriteResponse(rw, http.StatusOK, p.ToDto())
}

/*
* Update product
 */
func UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	uuid := requestVars["product_uuid"]
	var request dto.ProductRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.Error())
		return
	}
	u, err := domain.UpdateProduct(uuid, request)
	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		return
	}

	WriteResponse(rw, http.StatusOK, u.ToDto())
}

/*
* Delete product
 */
func DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	uuid := requestVars["product_uuid"]

	u, err := domain.DeleteProduct(uuid)

	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.AsMessage())
		return
	}

	WriteResponse(rw, http.StatusOK, u.ToDto())
}

/*
* Get all products paginated
 */
func GetProducts(rw http.ResponseWriter, r *http.Request) {
	response := domain.GetAllProducts(r)
	WriteResponse(rw, http.StatusOK, response)
}

/*
* Search products
 */
func Search(rw http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	var response []dto.SingleProduct
	c := domain.Search(sanitize.Scripts(q))

	for _, cs := range c {
		response = append(response, cs.ToDto())
	}

	WriteResponse(rw, http.StatusOK, response)
}
