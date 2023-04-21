package view

import (
	"base/domain/entities"
	"base/domain/usecases/authorization"
	"base/view/http_error"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"io"
	"log"
	"net/http"
)

type newHTTPAuthorizationModule struct {
	useCases authorization.UseCases
}

func NewHTTPAuthorization(useCases authorization.UseCases) HttpModule {
	return &newHTTPAuthorizationModule{
		useCases: useCases,
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

	err = n.useCases.Login(ctx, credentials)
	if err != nil {
		log.Println("[login] Error Login", err)
		http_error.HandleError(w, err)
		return
	}

	// TODO: store secret key in a safe place.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
	tokenString, err := token.SignedString("super_secret")
	if err != nil {
		log.Println("[CreateJWT] Error SignedString", err)
		http_error.HandleError(w, err)
		return
	}

	secureCookie := securecookie.New([]byte("secret_cookie_key"), nil)
	encodedToken, err := secureCookie.Encode("cookie", tokenString)
	if err != nil {
		log.Println("[CreateJWT] Error Encode", err)
		http_error.HandleError(w, err)
		return
	}
	cookie := &http.Cookie{
		Name:     "cookie",
		Value:    encodedToken,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	_, err = w.Write([]byte(encodedToken))
	if err != nil {
		log.Println("[login] Error Write", err)
		http_error.HandleError(w, err)
		return
	}
}
