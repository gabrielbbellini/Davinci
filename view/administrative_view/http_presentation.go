package administrative_view

import (
	presentation_mod "davinci/domain/administrative_usecases/presentation"
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
	useCases presentation_mod.UseCases
	settings settings.Settings
}

func NewHTTPPresentationModule(settings settings.Settings, cases presentation_mod.UseCases) view.HttpModule {
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

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("[create] Error ReadAll", err)
		http_error.HandleError(w, err)
		return
	}

	var presentation entities.Presentation
	err = json.Unmarshal(b, &presentation)
	if err != nil {
		log.Println("[Create] Error Unmarshal", err)
		http_error.HandleError(w, http_error.NewBadRequestError("Apresentação inválida."))
		return
	}

	presentationId, err := n.useCases.Create(ctx, presentation, user.Id)
	if err != nil {
		log.Println("[Create] Error Create", err)
		http_error.HandleError(w, err)
		return
	}

	_, err = w.Write([]byte(strconv.FormatInt(presentationId, 10)))
	if err != nil {
		log.Println("[Create] Error Write", err)
		http_error.HandleError(w, err)
		return
	}
}

func (n newHTTPPresentationModule) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.User)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("[update] Error ReadAll", err)
		http_error.HandleError(w, err)
		return
	}

	var presentation entities.Presentation
	err = json.Unmarshal(b, &presentation)
	if err != nil {
		log.Println("[update] Error Unmarshal", err)
		http_error.HandleError(w, http_error.NewBadRequestError("Apresentação inválida."))
		return
	}

	presentationId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[update] Error ParseInt", err)
		http_error.HandleError(w, err)
		return
	}

	err = n.useCases.Update(ctx, presentationId, presentation, user.Id)
	if err != nil {
		log.Println("[Update] Error Update", err)
		http_error.HandleError(w, err)
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Println("[Update] Error Write", err)
		http_error.HandleError(w, err)
		return
	}
}

func (n newHTTPPresentationModule) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.User)

	presentationId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[delete] Error ParseInt", err)
		http_error.HandleError(w, err)
		return
	}

	err = n.useCases.Delete(ctx, presentationId, user.Id)
	if err != nil {
		log.Println("[delete] Error Delete", err)
		http_error.HandleError(w, err)
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Println("[delete] Error Write", err)
		http_error.HandleError(w, err)
		return
	}
}

func (n newHTTPPresentationModule) getAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(entities.User)

	presentations, err := n.useCases.GetAll(ctx, user.Id)
	if err != nil {
		log.Println("[getAll] Error GetAll", err)
		http_error.HandleError(w, err)
		return
	}

	b, err := json.Marshal(presentations)
	if err != nil {
		log.Println("[getAll] Error Marshal", err)
		http_error.HandleError(w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		log.Println("[getAll] Error Write", err)
		http_error.HandleError(w, err)
		return
	}
}

func (n newHTTPPresentationModule) getById(w http.ResponseWriter, r *http.Request) {
	presentationId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[getById] Error ParseInt", err)
		http_error.HandleError(w, err)
		return
	}

	ctx := r.Context()
	user := ctx.Value("user").(entities.User)
	presentations, err := n.useCases.GetById(ctx, presentationId, user.Id)
	if err != nil {
		log.Println("[getById] Error GetById", err)
		http_error.HandleError(w, err)
		return
	}

	b, err := json.Marshal(presentations)
	if err != nil {
		log.Println("[getById] Error Marshal", err)
		http_error.HandleError(w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		log.Println("[getById] Error Write", err)
		http_error.HandleError(w, err)
		return
	}
}
