// this is the main entry point of the application for API

package main

import (
	"log"
	"net/http"
	"salary-calculator/api"
)

func main() {
	http.HandleFunc("/process", api.ProcessInputAPI)
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
