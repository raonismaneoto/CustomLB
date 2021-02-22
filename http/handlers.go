package http

import (
	"encoding/json"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/gorilla/mux"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/raonismaneoto/CustomLB/lb"
	"io/ioutil"
	"log"
	"net/http"
)

func (a *Api) join(w http.ResponseWriter, r *http.Request) {
	var (
		node lb.Node
	)

	if err := json.NewDecoder(r.Body).Decode(&node); err != nil {
		Write(w, http.StatusBadRequest, "invalid payload")
		return
	}

	u, err := uuid.NewV4()
	if err != nil {
		Write(w, http.StatusInternalServerError, "unable to generate uuid")
	}

	a.LB.AddNode(u.String(), &node)

	Write(w, 201, &map[string]string{"id": ""})
}

func (a *Api) serveRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Write(w, http.StatusBadRequest, "unable to parse body")
	}
	req := &lb.Request{
		r.Method,
		r.RequestURI,
		body,
		r.Header,
	}

	response, err := a.LB.ResolveRequest(req)
	if err != nil {
		log.Print("Got the following error: ", err.Error())
	}

	Write(w, response.StatusCode, response.Body)
}

func (a *Api) leave(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nodeId := params["id"]
	a.LB.RemoveNode(nodeId)
	Write(w, 200, "")
}

func (a *Api) update(w http.ResponseWriter, r *http.Request) {
	var (
		patch jsonpatch.Patch
	)

	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		Write(w, http.StatusBadRequest, "invalid payload")
		return
	}

	params := mux.Vars(r)
	nodeId := params["id"]

	err := a.LB.UpdateNode(nodeId, patch)
	if err != nil {
		Write(w, http.StatusNotFound, "node not found")
	}

	Write(w, 200, "")
}

func Write(w http.ResponseWriter, statusCode int, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(i); err != nil {
		log.Println("Error :" + err.Error())
	}
}
