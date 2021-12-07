package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// Serve - start API server and handle requests
func Serve() {
	router := mux.NewRouter()

	router.HandleFunc("/health/check", HealthCheck).Methods("GET")

	router.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE")
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Headers", "*")
			w.WriteHeader(http.StatusOK)
		})

	srv := &http.Server{
		Handler:      router,
		Addr:         ":1982",
		WriteTimeout: 300 * time.Second,
		ReadTimeout:  300 * time.Second,
	}

	log.Info("Start listening on :1982")

	log.Fatal(srv.ListenAndServe())
}
