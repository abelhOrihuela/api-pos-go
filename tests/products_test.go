//

package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"pos.com/app/dto"
)

func TestCreateProducts(t *testing.T) {
	product := dto.ProductRequest{
		Name:    "Coca cola",
		Price:   50.55,
		Barcode: "1003",
	}

	// request api
	writer := makeRequest("POST", "/products", product, false)

	// parse response
	var response dto.ProductResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "Coca cola", response.Name)
	assert.Equal(t, 50.55, response.Price)
	assert.Equal(t, "1003", response.Barcode)
}

func TestGetAllProducts(t *testing.T) {
	writer := makeRequest("GET", "/products", nil, false)

	// parse response
	var response []dto.ProductResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	// assertions
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.NotEmpty(t, response)
	assert.Equal(t, 3, len(response))

}

func TestSearchProducts(t *testing.T) {
	writer := makeRequest("GET", "/search?q=1001", nil, false)

	// parse response
	var response []dto.ProductResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	// assertions
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.NotEmpty(t, response)
	assert.Equal(t, 1, len(response))
}
