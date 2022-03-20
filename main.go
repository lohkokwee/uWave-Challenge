package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	service "github.com/lohkokwee/uwave_challenge/service"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("static")))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/busstop/{busStopId}", service.RetrieveStopDetails).Methods("GET")
	r.HandleFunc("/busline/{busLineId}", service.RetrieveLineDetails).Methods("GET")

	log.Println("--- Web server started ---")
	http.ListenAndServe(":80", r)
}
