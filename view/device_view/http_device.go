package device_view

import (
	"base/domain/device_usecases/device"
	"base/domain/entities"
	"base/view"
	"base/view/http_error"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
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
	router.HandleFunc("/devices/{id}", n.getById).Methods("GET")
	router.HandleFunc("/devices", n.create).Methods("POST")
	router.HandleFunc("/devices/{id}", n.update).Methods("PUT")
	router.HandleFunc("/devices/{id}", n.delete).Methods("DELETE")
}

func (n newHTTPDeviceModule) create(w http.ResponseWriter, r *http.Request) {
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

	err = n.useCases.Create(ctx, dev, user.Id)
	if err != nil {
		log.Println("[Create] Error", err)
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

	err = n.useCases.Delete(ctx, dev, user.Id)
	if err != nil {
		log.Println("[Update] Error", err)
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
