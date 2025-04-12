package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Dojeto/mcp-test/models"
	"github.com/Dojeto/mcp-test/services"
)

type TodoHandler struct {
	service *services.TodoService
}

func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.service.CreateTodo(&todo); err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *TodoHandler) GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetTodos()
	if err != nil {
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}