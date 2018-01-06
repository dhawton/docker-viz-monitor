package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	go initWorker()
	mux := http.NewServeMux()
	
	mux.HandleFunc("/nodes", nodeHandler)
	mux.HandleFunc("/services", serviceHandler)
	mux.HandleFunc("/tasks", taskHandler)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8081", handler)
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(tasks)
}

func nodeHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(nodes)
}

func serviceHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(services)
}
