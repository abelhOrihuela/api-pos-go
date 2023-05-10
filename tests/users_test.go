package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/morkid/paginate"
	"github.com/stretchr/testify/assert"
	"pos.com/app/domain"
	"pos.com/app/dto"
)

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
	var response dto.UserResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "Jonh Doe", response.Username)
	assert.Equal(t, "fake@hello.com", response.Email)
	assert.Equal(t, "cashier", response.Role)
	assert.Empty(t, response.DeletedAt)
}

func TestGetUser(t *testing.T) {
	user, err := domain.FindUserByEmail("fake@hello.com")

	// request api
	if err != nil {
		assert.Fail(t, "User not found")
	}

	writer := makeRequest("GET", "/api/pos/users/"+user.Uuid, nil, true)

	var response dto.UserResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.NotEmpty(t, response)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "fake@hello.com", response.Email)
}

func TestUpdateUser(t *testing.T) {

	u, err := domain.FindUserByEmail("fake@hello.com")

	// request api
	if err != nil {
		assert.Fail(t, "User not found")
	}

	user := dto.UserRequestUpdate{
		Username: "Jonh Doe Jr",
		Role:     "admin",
	}

	writer := makeRequest("PUT", "/api/pos/users/"+u.Uuid, user, true)

	var response dto.UserResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.NotEmpty(t, response)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "fake@hello.com", response.Email)
}

func TestDeleteUser(t *testing.T) {

	u, err := domain.FindUserByEmail("fake@hello.com")

	// request api
	if err != nil {
		assert.Fail(t, "User not found")
	}

	writer := makeRequest("DELETE", "/api/pos/users/"+u.Uuid, nil, true)

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
