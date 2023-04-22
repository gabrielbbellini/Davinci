package device_view

import (
	"base/domain/device_usecases/resolution"
	"base/view"
	"base/view/http_error"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type newHTTPResolutionModule struct {
	useCases resolution.UseCases
}

func NewHTTPResolutionModule(cases resolution.UseCases) view.HttpModule {
	return &newHTTPResolutionModule{
		useCases: cases,
	}
}

func (n newHTTPResolutionModule) Setup(router *mux.Router) {
	router.HandleFunc("/resolutions", n.getAll).Methods(http.MethodGet)
	router.HandleFunc("/resolutions/{id}", n.getById).Methods(http.MethodGet)
}

func (n newHTTPResolutionModule) getAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resolutions, err := n.useCases.GetAll(ctx)
	if err != nil {
		log.Println("[getAll] Error", err)
		http_error.HandleError(w, err)
		return
	}

	b, err := json.Marshal(resolutions)
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

func (n newHTTPResolutionModule) getById(w http.ResponseWriter, r *http.Request) {
	resolutionId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	ctx := r.Context()
	res, err := n.useCases.GetById(ctx, resolutionId)
	if err != nil {
		log.Println("[getById] Error", err)
		http_error.HandleError(w, err)
		return
	}

	b, err := json.Marshal(res)
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
