package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/morkid/paginate"
	"github.com/stretchr/testify/assert"
	"pos.com/app/dto"
)

var userResponse dto.UserResponse

func TestCreateUser(t *testing.T) {
	user := dto.UserRequest{
		Username: "Jonh Doe",
		Email:    "fake@hello.com",
		Password: "secret",
		Role:     "cashier",
	}

	// request api
	writer := makeRequest("POST", "/api/pos/users", user, true)

	// parse response
	json.Unmarshal(writer.Body.Bytes(), &userResponse)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "Jonh Doe", userResponse.Username)
	assert.Equal(t, "fake@hello.com", userResponse.Email)
	assert.Equal(t, "cashier", userResponse.Role)
	assert.Empty(t, userResponse.DeletedAt)
}

func TestGetUser(t *testing.T) {

	writer := makeRequest("GET", "/api/pos/users/"+userResponse.Uuid, nil, true)

	var response dto.UserResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.NotEmpty(t, response)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "fake@hello.com", response.Email)
}

func TestUpdateUser(t *testing.T) {

	user := dto.UserRequestUpdate{
		Username: "Jonh Doe Jr",
		Role:     "admin",
	}

	writer := makeRequest("PUT", "/api/pos/users/"+userResponse.Uuid, user, true)

	var response dto.UserResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.NotEmpty(t, response)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "fake@hello.com", response.Email)
}

func TestDeleteUser(t *testing.T) {

	writer := makeRequest("DELETE", "/api/pos/users/"+userResponse.Uuid, nil, true)

	var response dto.UserResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.NotEmpty(t, response)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.NotEmpty(t, response.DeletedAt)
}

func TestGetAllUsers(t *testing.T) {
	writer := makeRequest("GET", "/api/pos/users", nil, true)

	// parse response
	var response paginate.Page
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.NotEmpty(t, response)
	assert.NotEmpty(t, response.Items)
	assert.Equal(t, 2, int(response.Total))
	assert.Equal(t, http.StatusOK, writer.Code)
}
