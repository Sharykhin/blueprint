package main

import (
	"context"
	"encoding/json"
	"github.com/Sharykhin/blueprint/chapter7/meander"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	meander.APIKey = "AIzaSyDhzVMK5Yj-SnY2HgveUPNpia31rM1RHPY"
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	})

	server := http.Server{
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		sigs := make(chan os.Signal)
		signal.Notify(sigs, os.Interrupt)
		sig := <-sigs
		log.Printf("Server has been interrupted: %v", sig)
		err := server.Shutdown(context.Background())
		if err != nil {
			log.Fatalf("Could not gracefully shutdown server: %v", err)
		}
	}()

	log.Println("Server started on", ":8080")
	log.Fatal(server.ListenAndServe())
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) {
	publicData := make([]interface{}, len(data))

	for i, d := range data {
		publicData[i] = meander.Public(d)
	}

	json.NewEncoder(w).Encode(publicData)
}
