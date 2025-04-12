package storage

import (
	"context"
	"log"

	"github.com/Dojeto/mcp-test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoStorage struct {
	collection *mongo.Collection
}

func NewTodoStorage(db *mongo.Database) *TodoStorage {
	return &TodoStorage{
		collection: db.Collection("todos"),
	}
}

func (s *TodoStorage) Create(todo *models.Todo) error {
	_, err := s.collection.InsertOne(context.Background(), todo)
	if err != nil {
		log.Printf("Failed to insert todo: %v", err)
	}
	return err
}

func (s *TodoStorage) GetAll() ([]models.Todo, error) {
	cursor, err := s.collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("Failed to fetch todos: %v", err)
		return nil, err
	}
	var todos []models.Todo
	if err := cursor.All(context.Background(), &todos); err != nil {
		log.Printf("Failed to decode todos: %v", err)
		return nil, err
	}
	return todos, nil
}