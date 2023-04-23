package administrative_view

import (
	"davinci/domain/administrative_usecases/device"
	"davinci/domain/entities"
	"davinci/settings"
	"davinci/view"
	"davinci/view/http_error"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

type newHTTPDeviceModule struct {
	useCases device.UseCases
	settings settings.Settings
}

func NewHTTPDeviceModule(settings settings.Settings, useCases device.UseCases) view.HttpModule {
	return &newHTTPDeviceModule{
		useCases: useCases,
		settings: settings,
	}
}

func (n newHTTPDeviceModule) Setup(router *mux.Router) {
	router.HandleFunc("/devices", n.getAll).Methods(http.MethodGet)
	router.HandleFunc("/devices/{id}", n.getById).Methods(http.MethodGet)
	router.HandleFunc("/devices", n.create).Methods(http.MethodPost)
	router.HandleFunc("/devices/{id}", n.update).Methods(http.MethodPut)
	router.HandleFunc("/devices/{id}", n.delete).Methods(http.MethodDelete)
}

func (n newHTTPDeviceModule) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.User)

	read, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("[create] Error ReadAll", err)
		http_error.HandleError(w, err)
		return
	}

	var dev entities.Device
	err = json.Unmarshal(read, &dev)
	if err != nil {
		log.Println("[create] Error Unmarshal", err)
		http_error.HandleError(w, http_error.NewBadRequestError("dispositivo inválido"))
		return
	}

	err = n.useCases.Create(ctx, dev, user.Id)
	if err != nil {
		log.Println("[Create] Error", err)
		http_error.HandleError(w, err)
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Println("[create] Error Write", err)
		http_error.HandleError(w, err)
		return
	}
}

func (n newHTTPDeviceModule) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.User)
	request := r.Body
	read, err := io.ReadAll(request)
	if err != nil {
		return
	}

	var dev entities.Device
	err = json.Unmarshal(read, &dev)
	if err != nil {
		return
	}

	dev.Id, err = strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	err = n.useCases.Update(ctx, dev, user.Id)
	if err != nil {
		log.Println("[Update] Error", err)
		http_error.HandleError(w, err)
		return
	}
}

func (n newHTTPDeviceModule) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.User)

	deviceId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[delete] Error ParseInt")
		http_error.HandleError(w, err)
		return
	}

	err = n.useCases.Delete(ctx, deviceId, user.Id)
	if err != nil {
		log.Println("[delete] Error Delete", err)
		http_error.HandleError(w, err)
		return
	}
}

func (n newHTTPDeviceModule) getAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.User)
	devices, err := n.useCases.GetAll(ctx, user.Id)
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

func (n newHTTPDeviceModule) getById(w http.ResponseWriter, r *http.Request) {
	deviceId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	ctx := r.Context()
	user := ctx.Value("user").(entities.User)
	devices, err := n.useCases.GetById(ctx, deviceId, user.Id)
	if err != nil {
		log.Println("[getById] Error", err)
		http_error.HandleError(w, err)
		return
	}

	b, err := json.Marshal(devices)
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
