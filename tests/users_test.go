package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"pos.com/app/dto"
)

func TestCreateUser(t *testing.T) {
	user := dto.UserRequest{
		Username: "Jonh Doe",
		Email:    "jonh@hello.com",
		Password: "secret",
		Role:     "cashier",
	}

	// request api
	writer := makeRequest("POST", "/users", user, false)

	// parse response
	var response dto.UserResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "Jonh Doe", response.Username)
	assert.Equal(t, "jonh@hello.com", response.Email)
	assert.Equal(t, "cashier", response.Role)
}
