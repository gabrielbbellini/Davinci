package infrastructure

import (
	"base/domain/entities"
	authorization_usecases "base/domain/usecases/authorization"
	device_usecases "base/domain/usecases/device"
	resolution_usecases "base/domain/usecases/resolution"
	authorization_repository "base/infrastructure/repositories/authorization"
	device_repository "base/infrastructure/repositories/device"
	resolution_repository "base/infrastructure/repositories/resolution"
	"base/view"
	"base/view/http_error"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
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
	db, err := sql.Open("mysql", "root:root@tcp(devserver:3306)/davinci?parseTime=true")
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
	router.Use(rootMiddleware)

	authorizationRepository := authorization_repository.NewRepository(db)
	authorizationUseCases := authorization_usecases.NewUseCases(authorizationRepository)
	view.NewHTTPAuthorization(authorizationUseCases).Setup(router)

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(apiMiddleware)

	authorizationHelper := authorizationHelper{useCases: authorizationUseCases}
	apiRouter.Use(authorizationHelper.authorizationMiddleware)

	deviceRepository := device_repository.NewRepository(db)
	resolutionRepository := resolution_repository.NewResolutionRepository(db)
	deviceUseCases := device_usecases.NewUseCases(deviceRepository)
	resolutionUseCase := resolution_usecases.NewUseCases(resolutionRepository)

	view.NewHTTPResolutionModule(resolutionUseCase).Setup(apiRouter)
	view.NewHTTPDeviceModule(deviceUseCases).Setup(apiRouter)

	return nil
}

// rootMiddleware set the response content type for the api as json.
func rootMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Set the response content type for the api as json
		w.Header().Set("Content-Type", "application/json")

		//Set the origin to allow all.
		w.Header().Set("Access-Control-Allow-Origin", "*")

		//Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
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

type authorizationHelper struct {
	useCases authorization_usecases.UseCases
}

// authorizationMiddleware check if the user has the cookie with the token and if the token is valid.
func (a authorizationHelper) authorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Check if the user has the cookie with the token
		cookie, err := r.Cookie("cookie")
		if err != nil {
			if err == http.ErrNoCookie {
				//If the user doesn't have the cookie, return an error
				log.Println("[authorizationMiddleware] Error http.ErrNoCookie", err)
				http_error.HandleError(w, http_error.NewUnauthorizedError(err.Error()))
				return
			}
			//If there is an error, return an error
			log.Println("[authorizationMiddleware] Error r.Cookie", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError(err.Error()))
			return
		}

		token, err := getTokenFromCookie(cookie)
		if err != nil {
			log.Println("[authorizationMiddleware] Error getTokenFromCookie", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError(err.Error()))
			return
		}

		//Check if the token is valid
		if !token.Valid {
			//If the token is not valid, return an error
			log.Println("[authorizationMiddleware] Error !token.Valid", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Token inválido"))
			return
		}

		tokenData, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			log.Println("[authorizationMiddleware] Error !ok && !token.Valid", err)
			next.ServeHTTP(w, r)
			return
		}

		userString, ok := tokenData["user"]
		if !ok {
			log.Println("[authorizationMiddleware] Error !ok", err)
			next.ServeHTTP(w, r)
			return
		}

		var user entities.User
		err = json.Unmarshal([]byte(userString.(string)), &user)
		if err != nil {
			log.Println("[authorizationMiddleware] Error json.Unmarshal", err)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)

		//Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

// getTokenFromCookie get the token from the cookie.
func getTokenFromCookie(cookie *http.Cookie) (*jwt.Token, error) {
	secureCookie := securecookie.New([]byte(view.SecretJWTKey), nil)
	var tokenString string
	err := secureCookie.Decode("token", cookie.Value, &tokenString)
	if err != nil {
		log.Println("[login] Error Decode", err)
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Println("[login] token.Method.(*jwt.SigningMethodHMAC) !ok", err)
			return nil, errors.New("error parsing token")
		}
		return []byte(view.SecretJWTKey), nil
	})
	if err != nil {
		log.Println("[isCookieValid] Error parsing token", err)
		return nil, err
	}

	return token, nil
}
