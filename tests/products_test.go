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

func TestCreateProducts(t *testing.T) {
	product := dto.ProductRequest{
		Name:        "Coca cola",
		Price:       50.55,
		Barcode:     "1003",
		Description: "coca cola 600ml",
		Category:    1,
	}

	// request api/7
	writer := makeRequest("POST", "/api/pos/products", product, true)

	// parse response
	var response dto.Product
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "Coca cola", response.Name)
	assert.Equal(t, 50.55, response.Price)
	assert.Equal(t, "1003", response.Barcode)
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
