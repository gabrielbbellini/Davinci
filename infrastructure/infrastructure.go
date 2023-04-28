package infrastructure

import (
	"database/sql"
	"davinci/settings"
	"errors"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func Setup(settings settings.Settings, router *mux.Router) error {
	db, err := setupDataBase(settings)
	if err != nil {
		log.Println("[Setup] Error setupDataBase", err)
		return err
	}

	err = setupModules(settings, router, db)
	if err != nil {
		log.Println("[Setup] Error setupModules", err)
		return err
	}

	return nil
}

// SetupDataBase set the connection to the database and set connection entities.go.
func setupDataBase(settings settings.Settings) (*sql.DB, error) {
	db, err := sql.Open("mysql", settings.GetDBSource())
	if err != nil {
		log.Println("[Setup] Error connecting to database", err)
		return nil, err
	}

	// Limit the amount of time the connections are kept in the pool
	db.SetConnMaxLifetime(time.Minute * 10)

	// Limit the number of connections stored in the pool
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)

	return db, nil
}

// setupModules set the MVC structure for the application.
func setupModules(settings settings.Settings, router *mux.Router, db *sql.DB) error {
	router.Use(rootMiddleware)

	err := setupAdministrativeModules(settings, router, db)
	if err != nil {
		log.Println("[SetupMVC] Error setupAdministrativeModules", err)
		return err
	}

	err = setupDeviceModules(settings, router, db)
	if err != nil {
		log.Println("[SetupMVC] Error setupDeviceModules", err)
		return err
	}

	return nil
}

// rootMiddleware set the response content type for the api as json.
func rootMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Set the origin to allow all.
		w.Header().Set("Access-Control-Allow-Origin", "*")

		//Set the valid methods to all.
		w.Header().Set("Access-Control-Allow-Methods", "*")

		//Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// getTokenFromRequest get the token from the cookie.
func getTokenFromRequest(r *http.Request) (*jwt.Token, error) {
	splitToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if len(splitToken) != 2 {
		log.Println("[getTokenFromRequest] Error len(splitToken) == 0")
		return nil, errors.New("error Authorization Bearer not valid")
	}
	tokenString := splitToken[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Println("[getTokenFromRequest] token.Method.(*jwt.SigningMethodHMAC) !ok")
			return nil, errors.New("error parsing token")
		}
		return []byte(os.Getenv("DAVINCI_SECRET_KEY")), nil
	})
	if err != nil {
		log.Println("[getTokenFromRequest] Error parsing token", err)
		return nil, err
	}

	return token, nil
}
