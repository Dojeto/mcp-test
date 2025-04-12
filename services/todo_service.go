package services

import (
	"github.com/Dojeto/mcp-test/models"
	"github.com/Dojeto/mcp-test/storage"
)

type TodoService struct {
	storage *storage.TodoStorage
}

func NewTodoService(storage *storage.TodoStorage) *TodoService {
	return &TodoService{storage: storage}
}

func (s *TodoService) CreateTodo(todo *models.Todo) error {
	return s.storage.Create(todo)
}

func (s *TodoService) GetTodos() ([]models.Todo, error) {
	return s.storage.GetAll()
}