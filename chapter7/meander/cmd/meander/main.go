package main

import (
	"encoding/json"
	"github.com/Sharykhin/blueprint/chapter7/meander"
	"log"
	"net/http"
)

func main() {
	meander.APIKey = "AIzaSyDhzVMK5Yj-SnY2HgveUPNpia31rM1RHPY"
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	})

	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))

	for i, d := range data {
		publicData[i] = meander.Public(d)
	}

	return json.NewEncoder(w).Encode(publicData)
}
