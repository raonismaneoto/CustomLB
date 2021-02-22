package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Api struct {
	server *http.Server
}

func (a *Api) Start(port string) error {
	a.server= &http.Server{
		Addr:    ":" + port,
		Handler: a.handlers(),
	}
	log.Println("Starting api")
	return a.server.ListenAndServe()
}

func (a *Api) handlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/healthcheck", a.healthCheck).Methods(http.MethodGet)

	return router
}

func (a *Api) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

