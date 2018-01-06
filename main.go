package main

import (
	"net/http"
	"encoding/json"
)

func main() {
	go initWorker()

	http.HandleFunc("/tasks", taskHandler)
	http.ListenAndServe(":8081", nil)
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(tasks)
}