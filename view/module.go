package view

import "github.com/gorilla/mux"

type HttpModule interface {
	Setup(router *mux.Router)
}
