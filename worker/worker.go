package main

import (
	"encoding/json"
	"fmt"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"log"
	"time"
)

const (
	NodeEndpoint = "localhost:8081"
	LBEndpoint   = "localhost:8080"
)

type Worker struct {
	id            string
	UsedCpu       float32 // varies from 0 to 1
	UsedRam       float32 // varies from 0 to 1
	lastUpdatedAt time.Time
}

func (w *Worker) Start() {
	w.Join()
	go w.Update()
}

func (w *Worker) Join() {
	var reqBody map[string]interface{}

	payload, err := json.Marshal(w)
	if err != nil {
		fmt.Printf("error on marshalling payload")
		panic(w)
	}

	json.Unmarshal(payload, &reqBody)

	req := &Request{
		Body:     &reqBody,
		Endpoint: LBEndpoint + "/api/lb/nodes",
		Method:   "POST",
		Headers:  nil,
	}

	response, err := SendRequest(req)
	if err != nil {
		log.Fatal("unable to join LB")
	}

	var parsedResponse map[string]interface{}
	err = json.Unmarshal(response.Body, &parsedResponse)

	if id, ok := parsedResponse["id"]; ok {
		w.id = id.(string)
	}
}

func (w *Worker) Update() {
	for {
		oldWorker, err := json.Marshal(w)

		if err != nil {
			fmt.Printf("Error on parsing worker: ", err.Error())
		}

		w.Resources()

		currWorker, err := json.Marshal(w)

		if err != nil {
			fmt.Printf("Error on parsing worker: ", err.Error())
		}

		patch, err := jsonpatch.CreateMergePatch(oldWorker, currWorker)

		req := &Request{
			Body:     &patch,
			Endpoint: LBEndpoint + "/api/lb/nodes/" + w.id,
			Method:   "PATCH",
			Headers:  nil,
		}

		resp, err := SendRequest(req)
		if err != nil {
			log.Fatal("unable to update node")
		}

		if resp.StatusCode >= 300 {
			log.Fatal("Error on updating node with status code: %d", resp.StatusCode)
		}

		time.Sleep(60 * time.Second)
	}
}

func (w *Worker) Remove() {
	req := &Request{
		Body:     nil,
		Endpoint: "ap/lb/nodes/" + w.id,
		Method:   "DELETE",
		Headers:  nil,
	}

	resp, err := SendRequest(req)
	if err != nil {
		log.Fatal("unable to remove node")
	}

	if resp.StatusCode >= 300 {
		log.Fatal("Error on removing node with status code: %d", resp.StatusCode)
	}
}

func (w *Worker) Resources() {
	mem, err := memory.Get()

	if err != nil {
		log.Fatal("Unable to get ram info")
	}

	w.UsedRam = float32(mem.Used / mem.Total)

	cpu, err := cpu.Get()

	if err != nil {
		log.Fatal("Unable to get cpu info")
	}

	w.UsedCpu = float32(cpu.System / cpu.Total)
}
