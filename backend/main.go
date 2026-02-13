package main

import (
	"fmt"
	"log"
	"maze-game/api"
	"maze-game/store"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// "fmt"
// "net/http"

// Main entry point for the application.
// In Go, the execution starts with the 'main' function in the 'main' package.
func main() {
	// TODO: 1. Initialize your configuration or load environment variables if needed.
	// For example: port := os.Getenv("PORT")
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// TODO: 2. Load your data.
	// We need to load the questions from the JSON file at the start.
	// Call store.LoadQuestions() here.
	if err := store.LoadQuestions("store/questions.json"); err != nil {
		log.Fatalf("Failed to load questions: %v", err)
	}

	// TODO: 3. Set up your HTTP server and routes.
	// We will use the 'api' package to define our routes.
	// router := api.NewRouter()
	fmt.Println("Server starting on port", port)
	router := api.NewRouter()

	// TODO: 4. Start the server.
	// Use net/http to listen on a specific port (e.g., 8080).
	fmt.Println("Server starting on port", port, "...")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
