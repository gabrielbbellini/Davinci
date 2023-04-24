package device_view

import (
	"davinci/domain/device_usecases/presentation"
	"davinci/view"
	"davinci/view/http_error"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	router.HandleFunc("/presentation", n.getCurrentPresentation).Methods(http.MethodGet)
}

func (n newHTTPPresentationModule) getCurrentPresentation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	deviceId := ctx.Value("deviceId").(int64)

	presentations, err := n.useCases.GetCurrentPresentation(ctx, deviceId)
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
