package main

import (
	"encoding/json"
	"log"
	"net/http"
)
type M map[string]interface{}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := M{
		"message": "hello",
	}
	jsonData, err := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "ERROR:" + err.Error(), http.StatusInternalServerError)
	}
	_, _ = w.Write(jsonData)
}

func main() {
	mux := new(http.ServeMux)
	mux.HandleFunc("/", HomeHandler)

	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = mux

	log.Println("Server running at :9000")
	log.Fatal(server.ListenAndServe())
}
