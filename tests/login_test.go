package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"pos.com/app/dto"
)

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
