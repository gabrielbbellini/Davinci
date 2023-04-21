package infrastructure

import (
	device_usecases "base/domain/usecases/device"
	device_repository "base/infrastructure/repositories/device"
	"base/view"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func Setup(router *mux.Router) error {
	db, err := setupDataBase()
	if err != nil {
		log.Println("[Setup] Error setupDataBase", err)
		return err
	}

	err = setupMVC(router, db)
	if err != nil {
		log.Println("[Setup] Error setupMVC", err)
		return err
	}

	return nil
}

// SetupDataBase set the connection to the database and set connection settings.
func setupDataBase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/davinci")
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

// SetupMVC set the MVC structure for the application.
func setupMVC(router *mux.Router, db *sql.DB) error {
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(apiMiddleware)
	apiRouter.Use(authorizationMiddleware)

	deviceRepository := device_repository.NewDeviceRepository(db)
	deviceUseCases := device_usecases.NewUseCases(deviceRepository)

	view.NewHTTPDeviceModule(deviceUseCases).Setup(apiRouter)

	return nil
}

// apiMiddleware set the response content type for the api as json.
func apiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Set the response content type for the api as json
		w.Header().Set("Content-Type", "application/json")

		//Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// authorizationMiddleware check if the user has the cookie with the token and if the token is valid.
func authorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Check if the user has the cookie with the token
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				//If the user doesn't have the cookie, return an error
				http.Error(w, "No cookie", http.StatusUnauthorized)
				return
			}
			//If there is an error, return an error
			http.Error(w, "No cookie", http.StatusUnauthorized)
			return
		}
		//Check if the token is valid
		if !isTokenValid(cookie.Value) {
			//If the token is not valid, return an error
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		//Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)

	})
}
