package socket

import (
	"encoding/json"
	"log"
	"net/http"

	"maze-game/game"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all origins for now (development)
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WSMoveRequest struct {
	Type      string `json:"type"` // "move"
	Direction string `json:"direction"`
}

type WSMoveResponse struct {
	Type    string      `json:"type"` // "update" or "error"
	Payload interface{} `json:"payload"`
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	// Extract Game ID from query param ?id=...
	gameID := r.URL.Query().Get("id")
	if gameID == "" {
		log.Println("No game ID provided")
		return
	}

	gameInstance, exists := game.GetGame(gameID)
	if !exists {
		log.Println("Game not found:", gameID)
		return
	}

	log.Printf("Player connected to game %s", gameID)

	// Listen for messages
	for {
		// Read message as JSON
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		var req WSMoveRequest
		if err := json.Unmarshal(message, &req); err != nil {
			log.Println("JSON error:", err)
			continue
		}

		if req.Type == "move" {
			// Process move
			result, _, err := game.MovePlayer(gameInstance, req.Direction)

			response := WSMoveResponse{
				Type: "update",
			}

			if err != nil {
				response.Type = "error"
				response.Payload = err.Error()
			} else {
				// OPTIMIZATION: Send only what changed (Player & Status).
				// The Board is static and large, sending it every 100ms kills performance.

				payload := map[string]interface{}{
					"result": result,
					// "game_state": gameInstance, // Too big!
					"player": gameInstance.Player,
					"status": gameInstance.Status,
				}
				response.Payload = payload
			}

			// Write response
			responseBytes, _ := json.Marshal(response)
			if err := ws.WriteMessage(websocket.TextMessage, responseBytes); err != nil {
				log.Println("Write error:", err)
				break
			}
		}
	}
}
