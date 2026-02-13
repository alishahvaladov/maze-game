package game

import (
	"maze-game/models"

	"github.com/google/uuid"
)

// In-Memory Store for active games.
// Since we aren't using a database, we keep games in a global map.

var activeGames = make(map[string]*models.GameState)

// NewGame creates a new game session.
func NewGame(rows, cols int) *models.GameState {
	// Generate a new board
	board := GenerateMaze(rows, cols)

	// Create player at (0,0)
	player := models.Player{
		CurrentPos: models.Position{X: 0, Y: 0},
		Lives:      3,
		Score:      0,
	}

	// Create GameState
	id := uuid.New().String()
	gameState := &models.GameState{
		ID:     id,
		Board:  board,
		Player: player,
		Status: "ACTIVE",
	}

	// Store in map
	activeGames[id] = gameState

	return gameState
}

// Retrieve a game by ID.
func GetGame(id string) (*models.GameState, bool) {
	// TODO: Check if id exists in activeGames.
	game, exists := activeGames[id]
	return game, exists
}
