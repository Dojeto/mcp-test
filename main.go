package main

import (
	"log"
	"net/http"

	"github.com/Dojeto/mcp-test/handlers"
	"github.com/Dojeto/mcp-test/services"
	"github.com/Dojeto/mcp-test/storage"
	"github.com/Dojeto/mcp-test/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// Initialize MongoDB client
	client := utils.GetMongoClient()
	defer func() {
		if err := client.Disconnect(nil); err != nil {
			log.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

	// Initialize database and storage
	db := client.Database("todo_app")
	storage := storage.NewTodoStorage(db)

	// Initialize services and handlers
	service := services.NewTodoService(storage)
	handler := handlers.NewTodoHandler(service)

	// Define routes
	http.HandleFunc("/todos", handler.GetTodosHandler)
	http.HandleFunc("/todos/create", handler.CreateTodoHandler)

	// Start the server
	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}