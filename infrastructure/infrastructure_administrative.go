package infrastructure

import (
	authorization_usecases "base/domain/administrative_usecases/authorization"
	device_usecases "base/domain/administrative_usecases/device"
	"base/domain/entities"
	authorization_repository "base/infrastructure/administrative_repository/authorization"
	device_repository "base/infrastructure/administrative_repository/device"
	"base/infrastructure/administrative_repository/resolution"
	"base/view/administrative_view"
	"base/view/http_error"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupAdministrativeModules(router *mux.Router, db *sql.DB) error {

	authorizationRepository := authorization_repository.NewRepository(db)
	resolutionRepository := resolution.NewResolutionRepository(db)
	deviceRepository := device_repository.NewRepository(db)

	authorizationUseCases := authorization_usecases.NewUseCases(authorizationRepository)
	deviceUseCases := device_usecases.NewUseCases(deviceRepository, resolutionRepository)

	administrativeRouter := router.PathPrefix("/administrative").Subrouter()
	administrativeRouter.Use(authorizationMiddleware)

	administrative_view.NewHTTPAuthorization(authorizationUseCases).Setup(administrativeRouter)
	administrative_view.NewHTTPDeviceModule(deviceUseCases).Setup(administrativeRouter)

	return nil
}

// SetupMiddleware check if the user has the cookie with the token and if the token is valid.
func authorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pathTemplate, err := mux.CurrentRoute(r).GetPathTemplate()
		if err != nil {
			log.Println("[authorizationMiddleware] Error", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Token inválido"))
			return
		}

		if pathTemplate == "/administrative/login" {
			next.ServeHTTP(w, r)
			return
		}

		//Check if the user has the cookie with the token
		cookie, err := r.Cookie("cookie")
		if err != nil {
			if err == http.ErrNoCookie {
				//If the user doesn't have the cookie, return an error
				log.Println("[authorizationMiddleware] Error http.ErrNoCookie", err)
				http_error.HandleError(w, http_error.NewUnauthorizedError("Token inválido"))
				return
			}
			//If there is an error, return an error
			log.Println("[authorizationMiddleware] Error r.Cookie", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Token inválido"))
			return
		}

		token, err := getTokenFromCookie(cookie)
		if err != nil {
			log.Println("[authorizationMiddleware] Error getTokenFromCookie", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Token inválido"))
			return
		}

		//Check if the token is valid
		if !token.Valid {
			//If the token is not valid, return an error
			log.Println("[authorizationMiddleware] Error !token.Valid", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Token inválido"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			log.Println("[authorizationMiddleware] Error !ok && !token.Valid", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Token inválido"))
			return
		}

		userString, ok := claims["user"]
		if !ok {
			log.Println("[authorizationMiddleware] Error !ok", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Token inválido"))
			return
		}

		var user entities.User
		err = json.Unmarshal([]byte(userString.(string)), &user)
		if err != nil {
			log.Println("[authorizationMiddleware] Error json.Unmarshal", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Token inválido"))
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)

		//Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
