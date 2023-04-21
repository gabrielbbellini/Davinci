package device_view

import (
	"base/domain/device_usecases/device"
	"base/view"
	"base/view/http_error"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type newHTTPDeviceModule struct {
	useCases device.UseCases
}

func NewHTTPDeviceModule(cases device.UseCases) view.HttpModule {
	return &newHTTPDeviceModule{
		useCases: cases,
	}
}

func (n newHTTPDeviceModule) Setup(router *mux.Router) {
	router.HandleFunc("/devices", n.getAll).Methods("GET")
}

func (n newHTTPDeviceModule) getAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	devices, err := n.useCases.GetAll(ctx)
	if err != nil {
		log.Println("[getAll] Error", err)
		http_error.HandleError(w, err)
		return
	}

	b, err := json.Marshal(devices)
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
