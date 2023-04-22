package device_view

import (
	"base/domain/device_usecases/presentation"
	"base/domain/entities"
	"base/view"
	"base/view/http_error"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type newHTTPPresentationModule struct {
	useCases presentation.UseCases
}

func NewHTTPPresentationModule(cases presentation.UseCases) view.HttpModule {
	return &newHTTPPresentationModule{
		useCases: cases,
	}
}

func (n newHTTPPresentationModule) Setup(router *mux.Router) {
	router.HandleFunc("/presentations", n.getById).Methods("GET")
}

func (n newHTTPPresentationModule) getAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.User)
	resolution := ctx.Value("resolution").(entities.Resolution)
	presentations, err := n.useCases.GetAll(ctx, user.Id, resolution.Id)
	if err != nil {
		log.Println("[getAll] Error", err)
		http_error.HandleError(w, err)
		return
	}

	b, err := json.Marshal(presentations)
	if err != nil {
		log.Println("[getAll] Error Marshal", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		log.Println("[getAll] Error Write", err)
		return
	}
}

func (n newHTTPPresentationModule) getById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.User)
	device := ctx.Value("device").(entities.Device)
	presentations, err := n.useCases.GetById(ctx, device.Id, user.Id)
	if err != nil {
		log.Println("[getById] Error", err)
		http_error.HandleError(w, err)
		return
	}

	b, err := json.Marshal(presentations)
	if err != nil {
		log.Println("[getById] Error Marshal", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		log.Println("[getById] Error Write", err)
		return
	}
}
