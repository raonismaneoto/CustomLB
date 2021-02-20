package http

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Api struct {
	server *http.Server
}

func New() *Api {
	return &Api{}
}

func (a *Api) Start(port string) error {
	a.server= &http.Server{
		Addr:    ":" + port,
		Handler: a.handlers(),
	}
	log.Println("Starting worker api")
	return a.server.ListenAndServe()
}

func (a *Api) handlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/lb/nodes/join", a.join).Methods(http.MethodPost)
	router.HandleFunc("/api/lb/nodes", a.leave).Methods(http.MethodDelete)
	router.HandleFunc("/", a.serveRequest).Methods(http.MethodPost)

	return router
}
