//

package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/morkid/paginate"
	"github.com/stretchr/testify/assert"
	"pos.com/app/dto"
)

var productResponse dto.SingleProduct

func TestCreateProducts(t *testing.T) {
	product := dto.ProductRequest{
		Name:             "Coca cola",
		Price:            50.55,
		Barcode:          "10003",
		Description:      "coca cola 800ml",
		CategoryID:       1,
		Unit:             "PZA",
		CurrentExistence: 10,
	}

	// request api/7
	writer := makeRequest("POST", "/api/pos/products", product, true)

	// parse response
	json.Unmarshal(writer.Body.Bytes(), &productResponse)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "Coca cola", productResponse.Name)
	assert.Equal(t, 50.55, productResponse.Price)
	assert.Equal(t, "10003", productResponse.Barcode)
	assert.NotNil(t, productResponse.Id)
	assert.Equal(t, 1, productResponse.CategoryID)
}

func TestGetProduct(t *testing.T) {

	writer := makeRequest("GET", "/api/pos/products/"+productResponse.Uuid, nil, true)

	// parse response
	var response dto.SingleProduct
	json.Unmarshal(writer.Body.Bytes(), &response)

	// assertions
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, productResponse.Uuid, response.Uuid)
}

func TestUpdateProduct(t *testing.T) {
	productRequest := dto.ProductRequest{
		Name:             "Coca cola 600ml",
		Price:            65.50,
		Barcode:          "1005",
		Description:      "coca cola 600ml",
		CategoryID:       1,
		Unit:             "PZA",
		CurrentExistence: 100,
	}

	writer := makeRequest("PUT", "/api/pos/products/"+productResponse.Uuid, productRequest, true)

	// parse response
	var response dto.SingleProduct
	json.Unmarshal(writer.Body.Bytes(), &response)

	// assertions
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, productRequest.Name, response.Name)
	assert.Equal(t, productRequest.Description, response.Description)
	assert.Equal(t, productRequest.Price, response.Price)
	assert.Equal(t, productRequest.Barcode, response.Barcode)
	assert.NotNil(t, response.Id)
	assert.Equal(t, 1, response.CategoryID)
}

func TestDeleteProduct(t *testing.T) {

	writer := makeRequest("DELETE", "/api/pos/products/"+productResponse.Uuid, nil, true)

	// parse response
	var response dto.SingleProduct
	json.Unmarshal(writer.Body.Bytes(), &response)

	// assertions
	assert.NotEmpty(t, response)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.NotEmpty(t, response.DeletedAt)
}

func TestGetAllProducts(t *testing.T) {
	writer := makeRequest("GET", "/api/pos/products", nil, true)

	// parse response
	var response paginate.Page
	json.Unmarshal(writer.Body.Bytes(), &response)

	// assertions
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.NotEmpty(t, response.Items)
	assert.Equal(t, 2, int(response.Total))

}

func TestSearchProducts(t *testing.T) {
	writer := makeRequest("GET", "/api/pos/search?q=1002", nil, true)

	// parse response
	var response []dto.Product
	json.Unmarshal(writer.Body.Bytes(), &response)

	// assertions
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.NotEmpty(t, response)
	assert.Equal(t, 1, len(response))
}
