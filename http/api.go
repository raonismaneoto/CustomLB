package http

import (
	"github.com/raonismaneoto/CustomLB/lb"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Api struct {
	server *http.Server
	LB *lb.LoadBalancer
}

func New() *Api {
	return &Api{LB: &lb.LoadBalancer{}}
}

func (a *Api) Start(port string) error {
	a.server= &http.Server{
		Addr:    ":" + port,
		Handler: a.handlers(),
	}
	log.Println("Starting api")
	a.LB.Start()
	return a.server.ListenAndServe()
}

func (a *Api) handlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/lb/nodes", a.join).Methods(http.MethodPost)
	router.HandleFunc("/api/lb/nodes/{id}", a.update).Methods(http.MethodPut)
	router.HandleFunc("/api/lb/nodes/{id}", a.leave).Methods(http.MethodDelete)
	router.HandleFunc("/", a.serveRequest).Methods(http.MethodPost)
	router.HandleFunc("/", a.serveRequest).Methods(http.MethodPut)
	router.HandleFunc("/", a.serveRequest).Methods(http.MethodDelete)
	router.HandleFunc("/", a.serveRequest).Methods(http.MethodGet)
	router.HandleFunc("/", a.serveRequest).Methods(http.MethodHead)

	return router
}
