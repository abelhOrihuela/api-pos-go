//

package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"pos.com/app/api"
	"pos.com/app/db"
	"pos.com/app/domain"
	"pos.com/app/dto"
)

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	teardown()
	os.Exit(exitCode)
}

func TestHeartbeat(t *testing.T) {
	writer := makeRequest("GET", "/heartbeat", nil, false)

	// parse response
	var response map[string]interface{}
	json.Unmarshal(writer.Body.Bytes(), &response)

	// assertions
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "Online...", response["message"])

}

func TestGetRequest(t *testing.T) {
	writer := makeRequest("GET", "/products", nil, false)

	// parse response
	var response []dto.ProductResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	// assertions
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.NotEmpty(t, response)
	assert.Equal(t, 1, len(response))

}

func makeRequest(method, url string, body interface{}, isAuthenticatedRequest bool) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	if err != nil {
		panic(err)
	}

	if isAuthenticatedRequest {
		request.Header.Add("Authorization", "Bearer "+bearerToken())
	}

	writer := httptest.NewRecorder()
	api.Router().ServeHTTP(writer, request)
	return writer
}

func bearerToken() string {
	return ""
}

func setup() {
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()
	db.Database.AutoMigrate(&domain.Product{})
	db.Database.Create(&domain.Product{Name: "D42", Barcode: "100"})
}

func teardown() {
	migrator := db.Database.Migrator()
	migrator.DropTable(&domain.Product{})
	//migrator.DropTable(&model.Entry{})
}
