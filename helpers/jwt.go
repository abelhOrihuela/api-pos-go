package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"pos.com/app/db"
	"pos.com/app/domain"
	"pos.com/app/errs"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user *domain.User) (string, *errs.AppError) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Id,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	tokenRaw, err := token.SignedString(privateKey)

	if err != nil {
		return "", errs.NewDefaultError("Canot authenticate user")
	}

	return tokenRaw, nil
}

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if ValidateJWT(rw, r) {
			next.ServeHTTP(rw, r)
		} else {
			writeResponse(rw, http.StatusUnprocessableEntity, errs.NewDefaultError("Invalid token to auth"))
		}

	})
}

func ValidateJWT(rw http.ResponseWriter, r *http.Request) bool {
	token, err := getToken(r)

	if err != nil {
		return false
	} else {
		_, ok := token.Claims.(jwt.MapClaims)

		if ok && token.Valid {
			return true
		} else {
			return false
		}

	}

}

func CurrentUser(rw http.ResponseWriter, r *http.Request) (*domain.User, *errs.AppError) {

	if ValidateJWT(rw, r) {
		token, _ := getToken(r)
		claims, _ := token.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		user, err := domain.FindUserById(userId)
		if err != nil {
			return nil, err
		}
		return user, nil

	} else {
		return nil, errs.NewDefaultError("Invalid token")

	}

}

func getToken(r *http.Request) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func writeResponse(rw http.ResponseWriter, statusCode int, data interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		panic(err)
	}
}

func TenantMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var databaseName string
		var isProd bool

		if os.Getenv("ENV") == "production" {
			isProd = true
			domain := strings.Split(r.Host, ":")[0]
			subdomain := strings.Split(domain, ".")[0]
			databaseName = "pos_" + subdomain
		} else {
			databaseName = "test.db"
			isProd = false
		}

		SetupDatabase(databaseName, isProd)
		next.ServeHTTP(rw, r)

	})
}

func SetupDatabase(database string, isProd bool) {
	db.Connect(database, isProd)

	var settings domain.Settings

	err := db.Database.Last(&settings).Error

	if err != nil {
		db.Database.AutoMigrate(&domain.Settings{})
		db.Database.Create(&domain.Settings{
			VersionDB: "0",
			RootUser:  "NO",
		})

	} else {
		if isProd {
			fmt.Print("version db")
			fmt.Print(settings.VersionDB)
			fmt.Print(os.Getenv("DB_VERSION"))
			if settings.VersionDB != os.Getenv("DB_VERSION") {

				db.Database.AutoMigrate(&domain.Product{})
				db.Database.AutoMigrate(&domain.Order{})
				db.Database.AutoMigrate(&domain.OrderProduct{})
				db.Database.AutoMigrate(&domain.User{})
				db.Database.AutoMigrate(&domain.Category{})
				db.Database.AutoMigrate(&domain.Settings{})

				domain.UpdateVersion(os.Getenv("DB_VERSION"))
			}

			if settings.RootUser == "NO" {
				db.Database.Create(&domain.User{
					Username: "admin@hola.com",
					Email:    "admin@hola.com",
					Password: "secret",
					Role:     "admin",
				})
			}
		}
	}

}
