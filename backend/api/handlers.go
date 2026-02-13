package api

import (
	"encoding/json"
	"maze-game/game"
	"maze-game/models"
	"maze-game/store"
	"net/http"
)

// Request/Response Structs
type StartGameRequest struct {
	Rows int `json:"rows"`
	Cols int `json:"cols"`
}

// Handler for starting a new game.
// Endpoint: POST /api/game/start
func StartGameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req StartGameRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	// Default dimensions
	if req.Rows <= 0 {
		req.Rows = 10
	}
	if req.Cols <= 0 {
		req.Cols = 10
	}

	// Create a new game instance.
	newGame := game.NewGame(req.Rows, req.Cols)

	json.NewEncoder(w).Encode(newGame)
}

// Request/Response Structs
type MoveRequest struct {
	Direction string `json:"direction"`
}

type MoveResponse struct {
	Result    string            `json:"result"`
	GameState *models.GameState `json:"game_state"`
	Question  *models.Question  `json:"question,omitempty"`
}

// Handler for moving the player.
// Endpoint: POST /api/game/{id}/move
func MoveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. Extract the 'id' from the URL parameters.
	id := r.PathValue("id")

	// 2. Decode the request body
	var req MoveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 3. Retrieve the game state
	gameInstance, exists := game.GetGame(id)
	if !exists {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	// 4. Call the game logic to process the move.
	result, qIDReturned, err := game.MovePlayer(gameInstance, req.Direction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 5. Respond with the result.
	response := MoveResponse{
		Result:    result,
		GameState: gameInstance,
	}

	// Pick random question if needed
	if result == "QuestionFound" {
		var qID int = -1

		// QuestionFound -> Player is on it, correct ID should have been returned.
		qID = qIDReturned

		if qID >= 0 {
			// Find in bank
			for _, q := range store.QuestionBank {
				if q.ID == qID {
					response.Question = &q
					break
				}
			}
		}

		// Fallback
		if response.Question == nil && len(store.QuestionBank) > 0 {
			q := store.QuestionBank[0]
			response.Question = &q
		}
	}

	json.NewEncoder(w).Encode(response)
}

type AnswerRequest struct {
	QuestionID int    `json:"question_id"`
	Answer     string `json:"answer"`
}

// Handler for answering a question.
// Endpoint: POST /api/game/{id}/answer
func AnswerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	var req AnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	gameInstance, exists := game.GetGame(id)
	if !exists {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	result, err := game.AnswerQuestion(gameInstance, req.QuestionID, req.Answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}
