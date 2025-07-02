package main

import (
	"log"
	"net/http"

	"github.com/orayew2002/jun/internal/api"
	"github.com/orayew2002/jun/internal/task"
)

func main() {
	manager := task.NewManager()
	handler := api.NewHandler(manager)

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateTask(w, r)

		case http.MethodGet:
			handler.GetTask(w, r)

		case http.MethodDelete:
			handler.DeleteTask(w, r)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
