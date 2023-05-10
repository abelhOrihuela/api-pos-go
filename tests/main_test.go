//

package tests

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
	writer := makeRequest("GET", "/api/public/heartbeat", nil, false)

	// parse response
	var response map[string]interface{}
	json.Unmarshal(writer.Body.Bytes(), &response)

	// assertions
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "Online...", response["message"])

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

	response := authUser()

	return response.AccessToken
}

func authUser() dto.TokenResponse {
	user := dto.LoginRequest{
		Email:    "hola@robot.com",
		Password: "secret",
	}
	writer := makeRequest("POST", "/api/public/login", user, false)

	var response dto.TokenResponse
	json.Unmarshal(writer.Body.Bytes(), &response)
	return response
}

func setup() {
	err := godotenv.Load("./../.env.test")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()

	db.Database.AutoMigrate(&domain.Product{})
	db.Database.AutoMigrate(&domain.Order{})
	db.Database.AutoMigrate(&domain.OrderProduct{})
	db.Database.AutoMigrate(&domain.User{})
	db.Database.AutoMigrate(&domain.Category{})

	seedDatabase()
}

func teardown() {
	migrator := db.Database.Migrator()
	migrator.DropTable(&domain.Product{})
	migrator.DropTable(&domain.Order{})
	migrator.DropTable(&domain.OrderProduct{})
	migrator.DropTable(&domain.User{})
	migrator.DropTable(&domain.Category{})
}

func seedDatabase() {
	db.Database.Create(&domain.User{Username: "superuser", Email: "jonh@hello.com", Password: "secret", Role: "admin"})
	db.Database.Create(&domain.User{Username: "testing", Email: "hola@robot.com", Password: "secret", Role: "cashier"})
	db.Database.Create(&domain.Category{Name: "General", Description: "Genaral category for all products"})
	db.Database.Create(&domain.Product{Name: "New product", Barcode: "1001", Description: "New product", Price: 50.50, CategoryID: 1, Unit: "PZA"})
	db.Database.Create(&domain.Product{Name: "Old product", Barcode: "1002", Description: "Old product", Price: 50.60, CategoryID: 1, Unit: "CJA"})

}
