package device_view

import (
	"base/domain/device_usecases/authorization"
	"base/domain/entities"
	"base/view"
	"base/view/http_error"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

const SecretJWTKey = "secret"

type newHTTPAuthorizationModule struct {
	useCases authorization.UseCases
}

func NewHTTPAuthorization(useCases authorization.UseCases) view.HttpModule {
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

	device, err := n.useCases.Login(ctx, credentials)
	if err != nil {
		log.Println("[login] Error Login", err)
		http_error.HandleError(w, err)
		return
	}

	// TODO: store secret key in a safe place.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"deviceId": strconv.FormatInt(device.Id, 10),
	})

	tokenString, err := token.SignedString([]byte(SecretJWTKey))
	if err != nil {
		log.Println("[login] Error SignedString", err)
		http_error.HandleError(w, err)
		return
	}

	_, err = w.Write([]byte(tokenString))
	if err != nil {
		log.Println("[login] Error Write", err)
		http_error.HandleError(w, err)
		return
	}
}
