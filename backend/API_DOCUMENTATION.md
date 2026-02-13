# Go Maze Game Backend API

## Base URL
`http://localhost:3000`

## HTTP Endpoints

### 1. Start New Game
**POST** `/api/game/start`

Initializes a new game session with a generated maze.

**Request Body:**
```json
{
  "rows": 20,
  "cols": 20
}
```

**Response:**
Returns the initial `GameState` object, including the `id` required for WebSocket connection.

### 2. Answer Question
*(Note: Question mechanics are currently disabled in standard gameplay, but endpoint remains for compatibility if needed)*
**POST** `/api/game/{id}/answer`

Submits an answer to a question.

**Request Body:**
```json
{
  "question_id": 1,
  "answer": "London"
}
```

---

## WebSocket Endpoints

### 1. Game Implementation (Movement)
**URL** `/ws?id={GAME_ID}`

Connects to the game session for real-time updates.

**Client -> Server Messages:**

1. **Move Player**
   ```json
   {
     "type": "move",
     "direction": "UP"
   }
   ```
   *Values: "UP", "DOWN", "LEFT", "RIGHT"*

**Server -> Client Messages:**

1. **Game Update**
   ```json
   {
     "type": "update",
     "payload": {
       "result": "Moved",
       "player": {
         "current_pos": { "x": 1, "y": 2 },
         "lives": 3,
         "score": 0
       },
       "status": "ACTIVE"
     }
   }
   ```
   *Possible results: "Moved", "Blocked", "Win"*

2. **Error**
   ```json
   {
     "type": "error",
     "payload": "Game not found"
   }
   ```
