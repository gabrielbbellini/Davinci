package administrative_view

import (
	"base/domain/administrative_usecases/authorization"
	"base/domain/entities"
	"base/infrastructure"
	"base/view"
	"base/view/http_error"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"io"
	"log"
	"net/http"
)

const SecretJWTKey = "secret"

type newHTTPAuthorizationModule struct {
	useCases   authorization.UseCases
	middleware infrastructure.AdministrativeMiddleware
}

func NewHTTPAuthorization(useCases authorization.UseCases, middleware infrastructure.AdministrativeMiddleware) view.HttpModule {
	return &newHTTPAuthorizationModule{
		useCases:   useCases,
		middleware: middleware,
	}
}

func (n newHTTPAuthorizationModule) Setup(router *mux.Router) {
	router.HandleFunc("/login", n.login).Methods("POST")
}

func (n newHTTPAuthorizationModule) login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("[login] Error ReadAll", err)
		http_error.HandleError(w, err)
		return
	}

	var credentials entities.Credential
	if err = json.Unmarshal(b, &credentials); err != nil {
		log.Println("[login] Error Unmarshal", err)
		http_error.HandleError(w, err)
		return
	}

	user, err := n.useCases.Login(ctx, credentials)
	if err != nil {
		log.Println("[login] Error Login", err)
		http_error.HandleError(w, err)
		return
	}

	userByte, err := json.Marshal(*user)
	if err != nil {
		log.Println("[login] Error Marshal", err)
		http_error.HandleError(w, err)
		return
	}

	// TODO: store secret key in a safe place.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": string(userByte),
	})

	tokenString, err := token.SignedString([]byte(SecretJWTKey))
	if err != nil {
		log.Println("[login] Error SignedString", err)
		http_error.HandleError(w, err)
		return
	}

	if err != nil {
		log.Println("[login] Error Encode", err)
		http_error.HandleError(w, err)
		return
	}
	secureCookie := securecookie.New([]byte(SecretJWTKey), nil)
	encodedTokenString, err := secureCookie.Encode("token", tokenString)
	if err != nil {
		log.Println("[login] Error Encode", err)
		http_error.HandleError(w, err)
		return
	}
	cookie := &http.Cookie{
		Name:  "cookie",
		Value: encodedTokenString,
	}

	http.SetCookie(w, cookie)
	n.middleware.SetupMiddleware()
}
