package main

import (
	"fmt"
	"net/http"
  "github.com/go-chi/chi/v5"
	"encoding/json"
)

type ResponseData struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func main() {
	r := chi.NewRouter()

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		response := ResponseData{Status: 200, Message: "ok"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)
}
