package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"pos.com/app/dto"
)

func TestCreateOrder(t *testing.T) {
	var productsOrder []dto.OrderProductRequest

	productsOrder = append(productsOrder, dto.OrderProductRequest{
		IdProduct: 1,
		Price:     100,
		Quantity:  2,
	})

	// order request
	orderRequest := dto.OrderRequest{
		Total:      500.0,
		TotalItems: 1,
		Products:   productsOrder,
	}

	writer := makeRequest("POST", "/api/pos/orders", orderRequest, true)

	//parse response
	var response dto.Order
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, orderRequest.Total, response.Total)
	assert.Equal(t, orderRequest.TotalItems, len(response.OrderProducts))
}
