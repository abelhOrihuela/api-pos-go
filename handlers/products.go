package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"pos.com/app/db"
	"pos.com/app/domain"
	"pos.com/app/dto"
)

func GetAllProducts(rw http.ResponseWriter, r *http.Request) {
	var response []dto.ProductResponse

	c := make([]domain.Product, 0)

	db.Database.Find(&c) // find product with integer primary key

	for _, cs := range c {
		response = append(response, cs.ToDto())
	}

	fmt.Print(len(response))

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}
