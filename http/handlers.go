package http

import (
	"encoding/json"
	"fmt"
	"github.com/raonismaneoto/CustomLB/lb"
	"io/ioutil"
	"log"
	"net/http"
)

func (a *Api) join(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(r.RequestURI)
	Write(w, 200, "not implemented yet")
}

func (a *Api) serveRequest(w http.ResponseWriter, r *http.Request) {
	req := &lb.Request{
		r.Method,
		r.RequestURI,
		ioutil.ReadAll(r.Body),
		r.Header,
	}

	response, err := a.LB.ResolveRequest(req)
	if err != nil {
		log.Print("Got an error with code %s", response.StatusCode)
		log.Print(response.Body)
	}

	Write(w, response.StatusCode, response.Body)
}

func (a *Api) leave(w http.ResponseWriter, r *http.Request) {
	Write(w, 200, "not implemented yet")
}

func Write(w http.ResponseWriter, statusCode int, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(i); err != nil {
		log.Println("Error :" + err.Error())
	}
}
