package main

import (
	"context"
	"encoding/json"
	"github.com/Sharykhin/blueprint/chapter7/meander"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

func main() {
	meander.APIKey = "AIzaSyDhzVMK5Yj-SnY2HgveUPNpia31rM1RHPY"
	http.HandleFunc("/journeys", cors(func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	}))

	http.HandleFunc("/recommendations", cors(func(w http.ResponseWriter, r *http.Request) {
		q := &meander.Query{
			Journey: strings.Split(r.URL.Query().Get("journey"), "|"),
		}
		var err error
		q.Lat, err = strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		q.Lng, err = strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		q.Radius, err = strconv.Atoi(r.URL.Query().Get("radius"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		q.CostRangeStr = r.URL.Query().Get("cost")
		places := q.Run()
		respond(w, r, places)
	}))

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

func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		f(w, r)
	}
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) {
	publicData := make([]interface{}, len(data))

	for i, d := range data {
		publicData[i] = meander.Public(d)
	}

	json.NewEncoder(w).Encode(publicData)
}
