package main

import (
	"github.com/raonismaneoto/CustomLB/http"
	"log"
)

func main() {
	api := http.New()

	if err := api.Start("8080"); err != nil {
		log.Fatal(err.Error())
	}
}
