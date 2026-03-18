package handler

import (
	"encoding/json"
	"go-todo/internal/storage"
	"net/http"
	"strconv"
	"strings"
)

type TodoHandler struct {
	store *storage.MemoryStorage
}

func NewTodoHandler(store *storage.MemoryStorage) *TodoHandler {
	return &TodoHandler{store: store}
}

func (h *TodoHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/todos", h.handleTodos)
	mux.HandleFunc("/todos/", h.handleTodoByID)
}

func (h *TodoHandler) handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		todos := h.store.GetAll()
		json.NewEncoder(w).Encode(todos)
	case http.MethodPost:
		var req struct {
			Title string `json:"title"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		todo := h.store.Create(req.Title)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todo)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TodoHandler) handleTodoByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		var req struct {
			Title *string `json:"title"`
			Done  *bool   `json:"done"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		todo, err := h.store.Update(id, req.Title, req.Done)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(todo)
	case http.MethodDelete:
		if err := h.store.Delete(id); err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
