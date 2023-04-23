package administrative_view

import (
	"davinci/domain/administrative_usecases/presentation"
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

type newHTTPPresentationModule struct {
	useCases presentation.UseCases
	settings settings.Settings
}

func NewHTTPPresentationModule(settings settings.Settings, cases presentation.UseCases) view.HttpModule {
	return &newHTTPPresentationModule{
		useCases: cases,
		settings: settings,
	}
}

func (n newHTTPPresentationModule) Setup(router *mux.Router) {
	router.HandleFunc("/presentations", n.getAll).Methods(http.MethodGet)
	router.HandleFunc("/presentations/{id}", n.getById).Methods(http.MethodGet)
	router.HandleFunc("/presentations", n.create).Methods(http.MethodPost)
	router.HandleFunc("/presentations/{id}", n.update).Methods(http.MethodPut)
	router.HandleFunc("/presentations/{id}", n.delete).Methods(http.MethodDelete)
}

func (n newHTTPPresentationModule) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.User)
	request := r.Body
	read, err := io.ReadAll(request)
	if err != nil {
		return
	}

	var presentationRequest entities.Presentation
	err = json.Unmarshal(read, &presentationRequest)
	if err != nil {
		return
	}

	err = n.useCases.Create(ctx, presentationRequest, user.Id)
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

	var presentationRequest entities.Presentation
	err = json.Unmarshal(read, &presentationRequest)
	if err != nil {
		return
	}

	presentationRequest.Id, err = strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	err = n.useCases.Update(ctx, presentationRequest, user.Id)
	if err != nil {
		log.Println("[Update] Error", err)
		http_error.HandleError(w, err)
		return
	}
}

func (n newHTTPPresentationModule) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.User)

	var presentationRequest entities.Presentation

	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[Update] Error converting id", err)
		return
	}
	presentationRequest.Id = id

	err = n.useCases.Delete(ctx, presentationRequest, user.Id)
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
