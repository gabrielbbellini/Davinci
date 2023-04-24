package device_view

import (
	"davinci/domain/device_usecases/device"
	"davinci/view"
	"github.com/gorilla/mux"
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
}
