package administrative_view

import (
	"base/domain/administrative_usecases/presentation"
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

type newHTTPPresentationModule struct {
	useCases presentation.UseCases
}

func NewHTTPPresentationModule(cases presentation.UseCases) view.HttpModule {
	return &newHTTPPresentationModule{
		useCases: cases,
	}
}

func (n newHTTPPresentationModule) Setup(router *mux.Router) {
	router.HandleFunc("/presentations", n.getAll).Methods("GET")
	router.HandleFunc("/presentations/{id}", n.getById).Methods("GET")
	router.HandleFunc("/presentations", n.create).Methods("POST")
	router.HandleFunc("/presentations/{id}", n.update).Methods("PUT")
	router.HandleFunc("/presentations/{id}", n.delete).Methods("DELETE")
}

func (n newHTTPPresentationModule) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.User)
	request := r.Body
	read, err := io.ReadAll(request)
	if err != nil {
		return
	}

	var dev entities.Presentation
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

func (n newHTTPPresentationModule) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.User)
	request := r.Body
	read, err := io.ReadAll(request)
	if err != nil {
		return
	}

	var dev entities.Presentation
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

func (n newHTTPPresentationModule) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.User)
	request := r.Body
	read, err := io.ReadAll(request)
	if err != nil {
		return
	}

	var dev entities.Presentation
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

func (n newHTTPPresentationModule) getAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.User)
	presentations, err := n.useCases.GetAll(ctx, user.Id)
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
	presentationId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	ctx := r.Context()
	user := ctx.Value("user").(entities.User)
	presentations, err := n.useCases.GetById(ctx, presentationId, user.Id)
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
