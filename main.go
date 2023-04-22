package main

import (
	"base/infrastructure"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const ServerUrl = "10.0.11.140:8000"

func main() {
	router := mux.NewRouter()
	err := infrastructure.Setup(router)
	if err != nil {
		log.Println("[main] Error SetupDeviceModules", err)
		return
	}

	server := &http.Server{
		Handler:      router,
		Addr:         ServerUrl,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("[main] Server is running on", ServerUrl)
	log.Fatal(server.ListenAndServe())
}
