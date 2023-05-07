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
		Email:    "jonh@hello.com",
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
	assert.Equal(t, "jonh@hello.com", response.Email)
	assert.Equal(t, "cashier", response.Role)
}

func TestGetUser(t *testing.T) {
	user, err := domain.FindUserByEmail("jonh@hello.com")

	// request api
	if err != nil {
		assert.Fail(t, "User not found")
	}

	writer := makeRequest("GET", "/api/pos/users/"+user.Uuid, nil, true)

	var response dto.UserResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.NotEmpty(t, response)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "jonh@hello.com", response.Email)

}

func TestUpdateUser(t *testing.T) {

	u, err := domain.FindUserByEmail("jonh@hello.com")

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
	assert.Equal(t, "jonh@hello.com", response.Email)

}

func TestLogin(t *testing.T) {

	user := dto.LoginRequest{
		Email:    "jonh@hello.com",
		Password: "secret",
	}
	writer := makeRequest("POST", "/api/public/login", user, false)

	var response dto.TokenResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.NotEmpty(t, response.AccessToken)
	assert.Equal(t, http.StatusOK, writer.Code)

}

func TestMe(t *testing.T) {
	writer := makeRequest("GET", "/api/pos/me", nil, true)

	var response dto.UserResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.NotEmpty(t, response)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "hola@robot.com", response.Email)
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
