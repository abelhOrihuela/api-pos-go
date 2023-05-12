package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"pos.com/app/dto"
)

var responseCreatedOrder dto.Order

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
	json.Unmarshal(writer.Body.Bytes(), &responseCreatedOrder)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, orderRequest.Total, responseCreatedOrder.Total)
	assert.Equal(t, orderRequest.TotalItems, len(responseCreatedOrder.OrderProducts))
}

func TestGetOrder(t *testing.T) {
	var response dto.Order
	writer := makeRequest("GET", "/api/pos/orders/"+responseCreatedOrder.Uuid, nil, true)

	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, 500.0, responseCreatedOrder.Total)
	assert.Equal(t, 1, len(responseCreatedOrder.OrderProducts))
}
