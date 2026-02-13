# Go Maze Game Backend API

## Base URL
`http://localhost:3000`

## Endpoints

### 1. Start New Game
**POST** `/api/game/start`

Initializes a new game session with a generated maze.

**Request Body:**
```json
{
  "board": {
    "size": 8
  }
}
```
*Note: The current implementation implicitly expects the size nested inside a `board` object because it reuses the `GameState` struct for decoding.*

### 2. Move Player
**POST** `/api/game/{id}/move`

Moves the player in a specified direction.

**Request Body:**
```json
{
  "direction": "UP"
}
```
*Values: "UP", "DOWN", "LEFT", "RIGHT"*

**Response Example:**
```json
{
  "result": "Moved"
}
```
*Possible values for "result": "Moved", "Invalid Move", "Blocked", "QuestionFound", "QuestionWallHit", "Win"*

### 3. Answer Question
**POST** `/api/game/{id}/answer`

Submits an answer to a question.

**Request Body:**
```json
{
  "question_id": 1,
  "answer": "Paris"
}
```

**Response Example:**
```json
{
  "correct": true,
  "game_state": { ... }
}
```
