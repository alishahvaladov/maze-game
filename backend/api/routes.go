package api

import (
	"net/http"
)

// Middleware for CORS
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// NewRouter sets up the routes for our application.
func NewRouter() http.Handler {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Map URL paths to the handler functions we defined in handlers.go.
	mux.HandleFunc("POST /api/game/start", StartGameHandler)
	mux.HandleFunc("POST /api/game/{id}/move", MoveHandler)
	mux.HandleFunc("POST /api/game/{id}/answer", AnswerHandler)

	return EnableCORS(mux)
}
