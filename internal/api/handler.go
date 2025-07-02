package api

import (
	"encoding/json"
	"net/http"

	"github.com/orayew2002/jun/internal/task"
)

type Handler struct {
	manager *task.Manager
}

func NewHandler(m *task.Manager) *Handler {
	return &Handler{manager: m}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	t := h.manager.CreateTask(r.Context())
	writeJSON(w, http.StatusCreated, t)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id query param required", http.StatusBadRequest)
		return
	}

	t, err := h.manager.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id query param required", http.StatusBadRequest)
		return
	}

	err := h.manager.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
