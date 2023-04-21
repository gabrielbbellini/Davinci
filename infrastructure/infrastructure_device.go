package infrastructure

import (
	authorization_usecases "base/domain/device_usecases/authorization"
	device_usecases "base/domain/device_usecases/device"
	authorization_repository "base/infrastructure/device_repository/authorization"
	device_repository "base/infrastructure/device_repository/device"
	"base/view/device_view"
	"base/view/http_error"
	"context"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func SetupDeviceModules(router *mux.Router, db *sql.DB) error {
	deviceRouter := router.PathPrefix("/device").Subrouter()
	deviceRouter.Use(deviceAuthorizationMiddleware)

	authorizationRepository := authorization_repository.NewRepository(db)
	authorizationUseCases := authorization_usecases.NewUseCases(authorizationRepository)
	device_view.NewHTTPAuthorization(authorizationUseCases).Setup(deviceRouter)

	apiRouter := deviceRouter.PathPrefix("/api").Subrouter()

	deviceRepository := device_repository.NewRepository(db)
	deviceUseCases := device_usecases.NewUseCases(deviceRepository)
	device_view.NewHTTPDeviceModule(deviceUseCases).Setup(apiRouter)

	return nil
}

// deviceAuthorizationMiddleware check if the user has the cookie with the token and if the token is valid.
func deviceAuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Check if the user has the cookie with the token
		cookie, err := r.Cookie("cookie")
		if err != nil {
			if err == http.ErrNoCookie {
				//If the user doesn't have the cookie, return an error
				log.Println("[deviceAuthorizationMiddleware] Error http.ErrNoCookie", err)
				http_error.HandleError(w, http_error.NewUnauthorizedError(err.Error()))
				return
			}
			//If there is an error, return an error
			log.Println("[deviceAuthorizationMiddleware] Error r.Cookie", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError(err.Error()))
			return
		}

		token, err := getTokenFromCookie(cookie)
		if err != nil {
			log.Println("[deviceAuthorizationMiddleware] Error getTokenFromCookie", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError(err.Error()))
			return
		}

		//Check if the token is valid
		if !token.Valid {
			//If the token is not valid, return an error
			log.Println("[deviceAuthorizationMiddleware] Error !token.Valid", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Token inválido"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			log.Println("[deviceAuthorizationMiddleware] Error !ok && !token.Valid", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Token inválido"))
			return
		}

		deviceIdString, ok := claims["deviceId"]
		if !ok {
			log.Println("[deviceAuthorizationMiddleware] Error !ok", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Token inválido"))
			return
		}

		deviceId, err := strconv.Atoi(deviceIdString.(string))
		if err != nil {
			log.Println("[deviceAuthorizationMiddleware] Error strconv.Atoi", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Token inválido"))
			return
		}

		ctx := context.WithValue(r.Context(), "deviceId", deviceId)

		//Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
