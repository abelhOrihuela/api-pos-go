//

package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/morkid/paginate"
	"github.com/stretchr/testify/assert"
	"pos.com/app/db"
	"pos.com/app/domain"
	"pos.com/app/dto"
)

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
	var response dto.Product
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "Coca cola", response.Name)
	assert.Equal(t, 50.55, response.Price)
	assert.Equal(t, "10003", response.Barcode)
	assert.NotNil(t, response.Id)
	assert.Equal(t, 1, response.CategoryID)
}

func TestGetAllProducts(t *testing.T) {
	writer := makeRequest("GET", "/api/pos/products", nil, true)

	// parse response
	var response paginate.Page
	json.Unmarshal(writer.Body.Bytes(), &response)

	// assertions
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.NotEmpty(t, response.Items)
	assert.Equal(t, 3, int(response.Total))

}

func TestSearchProducts(t *testing.T) {
	writer := makeRequest("GET", "/api/pos/search?q=1001", nil, true)

	// parse response
	var response []dto.Product
	json.Unmarshal(writer.Body.Bytes(), &response)

	// assertions
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.NotEmpty(t, response)
	assert.Equal(t, 1, len(response))
}

func TestGetProduct(t *testing.T) {
	var product domain.Product
	err := db.Database.Where("id = ?", 1).First(&product).Error

	if err != nil {
		assert.Fail(t, "Product not exist")
	}

	writer := makeRequest("GET", "/api/pos/products/"+product.Uuid, nil, true)

	// parse response
	var response dto.SingleProduct
	json.Unmarshal(writer.Body.Bytes(), &response)

	// assertions
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, product.Uuid, response.Uuid)
}

func TestUpdateProduct(t *testing.T) {
	var product domain.Product
	err := db.Database.Where("id = ?", 1).First(&product).Error

	if err != nil {
		assert.Fail(t, "Product not exist")
	}

	productRequest := dto.ProductRequest{
		Name:             "Coca cola 600ml",
		Price:            65.50,
		Barcode:          "1005",
		Description:      "coca cola 600ml",
		CategoryID:       1,
		Unit:             "PZA",
		CurrentExistence: 100,
	}

	writer := makeRequest("PUT", "/api/pos/products/"+product.Uuid, productRequest, true)

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
	var product domain.Product
	err := db.Database.Where("id = ?", 1).First(&product).Error

	if err != nil {
		assert.Fail(t, "Product not exist")
	}

	writer := makeRequest("DELETE", "/api/pos/products/"+product.Uuid, nil, true)

	// parse response
	var response dto.SingleProduct
	json.Unmarshal(writer.Body.Bytes(), &response)

	// assertions
	assert.NotEmpty(t, response)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.NotEmpty(t, response.DeletedAt)
}
