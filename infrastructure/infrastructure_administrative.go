package infrastructure

import (
	"context"
	"database/sql"
	authorization_usecases "davinci/domain/administrative_usecases/authorization"
	device_usecases "davinci/domain/administrative_usecases/device"
	presentation_usecase "davinci/domain/administrative_usecases/presentation"
	resolution_usecases "davinci/domain/administrative_usecases/resolution"
	"davinci/domain/entities"
	authorization_repository "davinci/infrastructure/administrative_repository/authorization"
	device_repository "davinci/infrastructure/administrative_repository/device"
	presentation_repository "davinci/infrastructure/administrative_repository/presentation"
	"davinci/infrastructure/administrative_repository/resolution"
	"davinci/settings"
	"davinci/view/administrative_view"
	"davinci/view/http_error"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setupAdministrativeModules(settings settings.Settings, router *mux.Router, db *sql.DB) error {
	authorizationRepository := authorization_repository.NewRepository(settings, db)
	authorizationUseCases := authorization_usecases.NewUseCases(settings, authorizationRepository)

	resolutionRepository := resolution.NewResolutionRepository(settings, db)
	resolutionUseCases := resolution_usecases.NewUseCases(settings, resolutionRepository)

	deviceRepository := device_repository.NewRepository(settings, db)
	deviceUseCases := device_usecases.NewUseCases(settings, deviceRepository, resolutionRepository)

	presentationRepository := presentation_repository.NewPresentationRepository(settings, db)
	presentationUseCase := presentation_usecase.NewUseCases(settings, resolutionRepository, presentationRepository)

	administrativeRouter := router.PathPrefix("/administrative").Subrouter()
	administrativeRouter.Use(authorizationMiddleware)

	administrative_view.NewHTTPAuthorization(settings, authorizationUseCases).Setup(administrativeRouter)
	administrative_view.NewHTTPDeviceModule(settings, deviceUseCases).Setup(administrativeRouter)
	administrative_view.NewHTTPResolutionModule(settings, resolutionUseCases).Setup(administrativeRouter)
	administrative_view.NewHTTPPresentationModule(settings, presentationUseCase).Setup(administrativeRouter)

	return nil
}

// SetupMiddleware check if the user has the cookie with the token and if the token is valid.
func authorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pathTemplate, err := mux.CurrentRoute(r).GetPathTemplate()
		if err != nil {
			log.Println("[authorizationMiddleware] Error", err)
			http_error.HandleError(w, http_error.NewForbiddenError("Credenciais inválidas."))
			return
		}

		if pathTemplate == "/administrative/login" {
			next.ServeHTTP(w, r)
			return
		}

		token, err := getTokenFromRequest(r)
		if err != nil {
			log.Println("[authorizationMiddleware] Error getTokenFromRequest", err)
			http_error.HandleError(w, http_error.NewForbiddenError("Credenciais inválidas."))
			return
		}

		if !token.Valid {
			log.Println("[authorizationMiddleware] Error !token.Valid", err)
			http_error.HandleError(w, http_error.NewForbiddenError("Credenciais inválidas."))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			log.Println("[authorizationMiddleware] Error !ok && !token.Valid", err)
			http_error.HandleError(w, http_error.NewForbiddenError("Credenciais inválidas."))
			return
		}

		userString, ok := claims["user"]
		if !ok {
			log.Println("[authorizationMiddleware] Error !ok", err)
			http_error.HandleError(w, http_error.NewForbiddenError("Credenciais inválidas."))
			return
		}

		var user entities.UserCredential
		err = json.Unmarshal([]byte(userString.(string)), &user)
		if err != nil {
			log.Println("[authorizationMiddleware] Error json.Unmarshal", err)
			http_error.HandleError(w, http_error.NewForbiddenError("Credenciais inválidas."))
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)

		//Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
