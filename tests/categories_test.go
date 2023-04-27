package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/morkid/paginate"
	"github.com/stretchr/testify/assert"
	"pos.com/app/dto"
)

func TestCreateCategory(t *testing.T) {
	category := dto.CategoryRequest{
		Name:        "Abarrotes",
		Description: "Abarrotes",
	}

	// request api/pos/categories
	writer := makeRequest("POST", "/api/pos/categories", category, true)

	// parse response
	var response dto.Category
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "Abarrotes", response.Name)
	assert.Equal(t, "Abarrotes", response.Description)
	assert.NotNil(t, response.Id)
	assert.NotNil(t, response.Uuid)
}

func TestGetAllCategories(t *testing.T) {

	// request api/pos/categories
	writer := makeRequest("GET", "/api/pos/categories", nil, true)

	// parse response
	var response paginate.Page
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, 2, int(response.Total))
}
