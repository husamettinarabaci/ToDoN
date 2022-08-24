package main

import (
	"net/http"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {

	wg.Add(1)
	go createHttpServer()
	wg.Wait()
}

// createHttpServer
// This function creates a new http-server and listens it
func createHttpServer() {
	http.HandleFunc("/health", HealthHandler)
	http.ListenAndServe(":80", nil)
}

// healthHandler
// This handler is used by K8S for probe of healty
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
