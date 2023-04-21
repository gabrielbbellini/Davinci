package view

import (
	"base/domain/entities"
	"base/domain/usecases/authorization"
	"base/view/http_error"
	"encoding/json"
	"github.com/gorilla/mux"
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

	_, err = w.Write([]byte("Success"))
	if err != nil {
		log.Println("[login] Error Write", err)
		http_error.HandleError(w, err)
		return
	}
}
