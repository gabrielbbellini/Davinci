package main

import (
	"davinci/infrastructure"
	configs "davinci/settings"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	settings, err := configs.Setup()
	if err != nil {
		log.Println("[main] Error configs.Setup", err)
		return
	}
	serverDomain := settings.GetDomain()

	router := mux.NewRouter()
	err = infrastructure.Setup(*settings, router)
	if err != nil {
		log.Println("[main] Error infrastructure.Setup", err)
		return
	}

	server := &http.Server{
		Handler: handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "Accept"}),
			handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}),
			handlers.AllowCredentials(),
		)(router),
		Addr:         serverDomain,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("[main] Server is running on", serverDomain)
	log.Fatal(server.ListenAndServe())
}
