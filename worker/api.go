package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Api struct {
	server *http.Server
	worker Worker
}

func (a *Api) Start(port string) error {
	a.server= &http.Server{
		Addr:    ":" + port,
		Handler: a.handlers(),
	}
	a.worker.Start()
	log.Println("Starting api")
	return a.server.ListenAndServe()
}

func (a *Api) handlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/healthcheck", a.healthCheck).Methods(http.MethodGet)
	router.HandleFunc("/api/resources", a.NodeResources).Methods(http.MethodGet)
	return router
}

func (a *Api) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (a *Api) NodeResources(w http.ResponseWriter, r *http.Request) {
	a.worker.Resources()
	response := &map[string]float32{"usedCpu": a.worker.UsedCpu, "usedRam": a.worker.UsedRam}
	Write(w, 200, response)
}

func Write(w http.ResponseWriter, statusCode int, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(i); err != nil {
		log.Println("Error :" + err.Error())
	}
}

