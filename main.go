package main

import (
	"net/http"
	"encoding/json"
)

func main() {
	go initWorker()

	http.HandleFunc("/tasks", taskHandler)
	http.HandleFunc("/nodes", nodeHandler)
	http.HandleFunc("/services", serviceHandler)
	http.ListenAndServe(":8081", nil)
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