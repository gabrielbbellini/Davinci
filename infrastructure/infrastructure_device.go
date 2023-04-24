package infrastructure

import (
	"context"
	"database/sql"
	authorization_usecases "davinci/domain/device_usecases/authorization"
	device_usecases "davinci/domain/device_usecases/device"
	presentation_usecases "davinci/domain/device_usecases/presentation"
	resolution_usecases "davinci/domain/device_usecases/resolution"
	"davinci/domain/entities"
	"encoding/json"

	authorization_repository "davinci/infrastructure/device_repository/authorization"
	device_repository "davinci/infrastructure/device_repository/device"
	presentation_repository "davinci/infrastructure/device_repository/presentation"
	resolution_repository "davinci/infrastructure/device_repository/resolution"
	user_repository "davinci/infrastructure/device_repository/user"

	"davinci/settings"
	"davinci/view/device_view"
	"davinci/view/http_error"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setupDeviceModules(settings settings.Settings, router *mux.Router, db *sql.DB) error {
	// TODO: PASS THE SETTINGS AS PARAMETER TO EACH USE CASE, REPOSITORY AND HTTP MODULE

	userRepository := user_repository.NewRepository(db)
	deviceRepository := device_repository.NewRepository(db)
	presentationRepository := presentation_repository.NewPresentationRepository(db)
	authorizationRepository := authorization_repository.NewRepository(db)
	resolutionRepository := resolution_repository.NewResolutionRepository(db)

	authorizationUseCases := authorization_usecases.NewUseCases(authorizationRepository, userRepository, deviceRepository)
	deviceUseCases := device_usecases.NewUseCases(deviceRepository)
	presentationUseCases := presentation_usecases.NewUseCases(presentationRepository)
	resolutionUseCases := resolution_usecases.NewUseCases(resolutionRepository)

	deviceRouter := router.PathPrefix("/device").Subrouter()
	device_view.NewHTTPAuthorization(authorizationUseCases).Setup(deviceRouter)

	deviceRouter.Use(deviceAuthorizationMiddleware)

	device_view.NewHTTPDeviceModule(deviceUseCases).Setup(deviceRouter)
	device_view.NewHTTPPresentationModule(presentationUseCases).Setup(deviceRouter)
	device_view.NewHTTPResolutionModule(resolutionUseCases).Setup(deviceRouter)

	return nil
}

// deviceAuthorizationMiddleware check if the user has the cookie with the token and if the token is valid.
func deviceAuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pathTemplate, err := mux.CurrentRoute(r).GetPathTemplate()
		if err != nil {
			log.Println("[deviceAuthorizationMiddleware] Error", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Credenciais inválidas."))
			return
		}

		if pathTemplate == "/device/login" {
			next.ServeHTTP(w, r)
			return
		}

		token, err := getTokenFromRequest(r)
		if err != nil {
			log.Println("[deviceAuthorizationMiddleware] Error getTokenFromRequest", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Credenciais inválidas."))
			return
		}

		//Check if the token is valid
		if !token.Valid {
			//If the token is not valid, return an error
			log.Println("[deviceAuthorizationMiddleware] Error !token.Valid", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Credenciais inválidas."))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			log.Println("[deviceAuthorizationMiddleware] Error !ok && !token.Valid", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Credenciais inválidas."))
			return
		}

		deviceString, ok := claims["device"]
		if !ok {
			log.Println("[deviceAuthorizationMiddleware] Error !ok", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Credenciais inválidas."))
			return
		}

		var device entities.Device
		err = json.Unmarshal([]byte(deviceString.(string)), &device)
		if err != nil {
			log.Println("[deviceAuthorizationMiddleware] Error strconv.Atoi", err)
			http_error.HandleError(w, http_error.NewUnauthorizedError("Credenciais inválidas."))
			return
		}

		ctx := context.WithValue(r.Context(), "device", device)

		//Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
