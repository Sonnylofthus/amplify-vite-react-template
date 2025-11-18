package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Todo represents a todo item
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

// Simple in-memory storage
var todos []Todo
var nextID = 1

func main() {
	// Initialize with some sample todos
	todos = []Todo{
		{ID: nextID, Title: "Learn Go", Completed: true, CreatedAt: time.Now()},
	}
	nextID++

	// Set up routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/todos", todosHandler)
	http.HandleFunc("/api/health", healthHandler)

	// Start server
	port := ":8080"
	fmt.Printf("🚀 Go web server starting on http://localhost%s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  / - Welcome message")
	fmt.Println("  GET  /api/todos - List all todos")
	fmt.Println("  POST /api/todos - Create a new todo (JSON body: {\"title\": \"...\"})")
	fmt.Println("  GET  /api/health - Health check")
	fmt.Println("\nPress Ctrl+C to stop the server")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// homeHandler handles the root route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome to the Go Learning API!",
		"version": "1.0.0",
		"docs":    "/api/todos for todo operations",
	})
}

// todosHandler handles todo operations
func todosHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS for React frontend
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// Handle preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case "GET":
		// Return all todos
		json.NewEncoder(w).Encode(todos)

	case "POST":
		// Create a new todo
		var newTodo struct {
			Title string `json:"title"`
		}

		if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if newTodo.Title == "" {
			http.Error(w, "Title is required", http.StatusBadRequest)
			return
		}

		todo := Todo{
			ID:        nextID,
			Title:     newTodo.Title,
			Completed: false,
			CreatedAt: time.Now(),
		}
		nextID++
		todos = append(todos, todo)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todo)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// healthHandler provides a health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now(),
		"todos":     len(todos),
	})
}
