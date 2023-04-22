package device_view

import (
	"base/domain/device_usecases/device"
	"base/view"
	"github.com/gorilla/mux"
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
	router.HandleFunc("/device/{id}", n.get).Methods("GET")
}

func (n newHTTPDeviceModule) get(w http.ResponseWriter, r *http.Request) {

}
