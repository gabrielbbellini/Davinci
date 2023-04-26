package administrative_view

import (
	"davinci/domain/administrative_usecases/device_presentation"
	"davinci/domain/entities"
	"davinci/settings"
	"davinci/view"
	"davinci/view/http_error"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type newHTTPDevicePresentationModule struct {
	useCases device_presentation.UseCases
	settings settings.Settings
}

func NewHTTPDevicePresentationModule(settings settings.Settings, useCases device_presentation.UseCases) view.HttpModule {
	return &newHTTPDevicePresentationModule{
		useCases: useCases,
		settings: settings,
	}
}

func (n newHTTPDevicePresentationModule) Setup(router *mux.Router) {
	router.HandleFunc("/devices/{deviceId}/presentations/{presentationId}", n.relate).Methods(http.MethodPost)
	router.HandleFunc("/devices/{deviceId}/presentations/current", n.getCurrentPresentation).Methods(http.MethodGet)
	router.HandleFunc("/devices/{deviceId}/presentations/current/{presentationId}", n.setCurrentPresentation).Methods(http.MethodPut)
}

func (n newHTTPDevicePresentationModule) relate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.UserCredential)

	deviceId, err := strconv.ParseInt(mux.Vars(r)["deviceId"], 10, 64)
	if err != nil {
		log.Println("[relate] Error ParseInt", err)
		http_error.HandleError(w, err)
		return
	}

	presentationId, err := strconv.ParseInt(mux.Vars(r)["presentationId"], 10, 64)
	if err != nil {
		log.Println("[relate] Error ParseInt", err)
		http_error.HandleError(w, err)
		return
	}

	err = n.useCases.Relate(ctx, user.Id, deviceId, presentationId)
	if err != nil {
		log.Println("[relate] Error", err)
		http_error.HandleError(w, err)
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		http_error.HandleError(w, err)
		log.Println("[relate] Error Write", err)
		return
	}
}

func (n newHTTPDevicePresentationModule) getCurrentPresentation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.UserCredential)

	deviceId, err := strconv.ParseInt(mux.Vars(r)["deviceId"], 10, 64)
	if err != nil {
		log.Println("[getCurrentPresentation] Error ParseInt", err)
		http_error.HandleError(w, err)
		return
	}

	presentations, err := n.useCases.GetCurrentPresentation(ctx, user.Id, deviceId)
	if err != nil {
		log.Println("[getCurrentPresentation] Error", err)
		http_error.HandleError(w, err)
		return
	}

	b, err := json.Marshal(presentations)
	if err != nil {
		log.Println("[getCurrentPresentation] Error Marshal", err)
		http_error.HandleError(w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		http_error.HandleError(w, err)
		log.Println("[getCurrentPresentation] Error Write", err)
		return
	}
}

func (n newHTTPDevicePresentationModule) setCurrentPresentation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.UserCredential)

	deviceId, err := strconv.ParseInt(mux.Vars(r)["deviceId"], 10, 64)
	if err != nil {
		log.Println("[setCurrentPresentation] Error ParseInt", err)
		http_error.HandleError(w, err)
		return
	}

	presentationId, err := strconv.ParseInt(mux.Vars(r)["presentationId"], 10, 64)
	if err != nil {
		log.Println("[setCurrentPresentation] Error ParseInt", err)
		http_error.HandleError(w, err)
		return
	}

	err = n.useCases.SetCurrentPresentation(ctx, user.Id, deviceId, presentationId)
	if err != nil {
		log.Println("[setCurrentPresentation] Error", err)
		http_error.HandleError(w, err)
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		http_error.HandleError(w, err)
		log.Println("[setCurrentPresentation] Error Write", err)
		return
	}
}
