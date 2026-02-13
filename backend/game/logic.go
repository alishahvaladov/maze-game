package game

import (
	"fmt"
	"math/rand"
	"maze-game/models"
	"maze-game/store"
	"time"
)

type AnswerResult struct {
	Correct   bool              `json:"correct"`
	GameState *models.GameState `json:"game_state"`
}

// Initialize the random seed.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Function to generate a new maze.
// Algorithm: Prim's Algorithm or Recursive Backtracking.
// Function to generate a new maze.
// Algorithm: Standard Recursive Backtracking to create a dense labyrinth.
func GenerateMaze(rows, cols int) models.Board {
	fmt.Println("Generating standard maze", rows, "x", cols)

	// Ensure odd dimensions for proper wall/path generation if using "jump 2" method.
	// adjustments might be needed if inputs are even, but let's try to handle it.
	// Actually, "jump 2" works best on odd grids.
	// If we get even constraints, we might have an extra wall edge, which is fine.

	// 1. Initialize grid with WALLS
	grid := make([][]models.Cell, rows)
	for y := 0; y < rows; y++ {
		grid[y] = make([]models.Cell, cols)
		for x := 0; x < cols; x++ {
			grid[y][x] = models.Cell{
				Type:     models.Wall,
				Position: models.Position{X: x, Y: y},
			}
		}
	}

	// 2. DFS Carving (Recursive Backtracking)
	// We carve paths.
	// Start at 0,0

	var carve func(cx, cy int)
	carve = func(cx, cy int) {
		grid[cy][cx].Type = models.Path

		// Directions: Up, Down, Left, Right (Jump 2 cells)
		dirs := [][]int{
			{0, -2, 0, -1}, // dx, dy, midX, midY relative
			{0, 2, 0, 1},
			{-2, 0, -1, 0},
			{2, 0, 1, 0},
		}
		rand.Shuffle(len(dirs), func(i, j int) { dirs[i], dirs[j] = dirs[j], dirs[i] })

		for _, d := range dirs {
			nx, ny := cx+d[0], cy+d[1]
			mx, my := cx+d[2], cy+d[3] // Wall between

			// Check bounds for target
			if nx >= 0 && nx < cols && ny >= 0 && ny < rows {
				if grid[ny][nx].Type == models.Wall {
					// Carve through wall
					grid[my][mx].Type = models.Path
					carve(nx, ny)
				}
			}
		}
	}

	// Start carving
	carve(0, 0)

	// 2.5. Braiding (Remove Dead Ends to create Loops)
	// "User must be confused too which way to go :/"
	// A perfect maze has no loops. Adding loops makes it harder (can't just follow walls).
	// Identify Dead Ends: Path cells with exactly 1 Path neighbor.

	// We iterate internally to find dead ends and open them up.
	// High braidingFactor = more loops = harder/more confusing.
	braidingFactor := 0.5 // 50% of dead ends turned into loops

	type Point struct{ X, Y int }

	// Helper to get path neighbors
	getPathNeighbors := func(cx, cy int) []Point {
		var neighbors []Point
		dirs := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // standard 1-step
		for _, d := range dirs {
			nx, ny := cx+d[0], cy+d[1]
			if nx >= 0 && nx < cols && ny >= 0 && ny < rows {
				if grid[ny][nx].Type == models.Path || grid[ny][nx].Type == models.Start {
					neighbors = append(neighbors, Point{X: nx, Y: ny})
				}
			}
		}
		return neighbors
	}

	for y := 1; y < rows-1; y++ {
		for x := 1; x < cols-1; x++ {
			if grid[y][x].Type == models.Path {
				neighbors := getPathNeighbors(x, y)
				if len(neighbors) == 1 {
					// DEAD END FOUND!
					if rand.Float64() < braidingFactor {
						// Try to connect to another path nearby (jump 2)
						// Check 4 directions for a wall that separates us from another path
						dirs := [][]int{
							{0, -2, 0, -1}, // jump, mid
							{0, 2, 0, 1},
							{-2, 0, -1, 0},
							{2, 0, 1, 0},
						}
						rand.Shuffle(len(dirs), func(i, j int) { dirs[i], dirs[j] = dirs[j], dirs[i] })

						connected := false
						for _, d := range dirs {
							nx, ny := x+d[0], y+d[1] // Target (potential path)
							mx, my := x+d[2], y+d[3] // Wall between

							if nx >= 0 && nx < cols && ny >= 0 && ny < rows {
								// If the target is ALSO a path (and not the one neighbor we come from), connect!
								// Actually, if it's a path, it's a valid candidate.
								if grid[ny][nx].Type == models.Path {
									grid[my][mx].Type = models.Path
									connected = true
									break
								}
							}
						}

						if !connected {
							// If we couldn't connect to a path, maybe just open a wall to valid space?
							// Usually braiding works best if we connect to existing paths.
							// If strictness fails, leave it as dead end.
						}
					}
				}
			}
		}
	}

	// Ensure Start and End are open (Backtracking usually guarantees coverage of reachable areas)
	// Force Start
	grid[0][0].Type = models.Start

	// Force End at bottom-right (or closest valid path cell?)
	// In odd-grid generation, usually the last cell might be a wall if dimensions are even.
	// Let's force proper End.
	endY, endX := rows-1, cols-1
	// If it's a wall, carve it or find neighbor?
	// Let's just force it to be END and carve path to it if needed?
	// Actually, simply setting it to END is enough, but we must ensure connectivity.
	// If the maze is perfect, (0,0) connects to everything carved.
	// Is (rows-1, cols-1) carved?
	// Only if rows/cols are odd.
	// If they are even, the bottom/right edges might be walls.
	// Quick fix: Carve a path from a neighbor if it's currently a Wall.

	if grid[endY][endX].Type == models.Wall {
		grid[endY][endX].Type = models.End
		// Check neighbors to connect
		connected := false
		// Try Up
		if endY > 0 && grid[endY-1][endX].Type != models.Wall {
			connected = true
		}
		// Try Left
		if !connected && endX > 0 && grid[endY][endX-1].Type != models.Wall {
			connected = true
		}

		if !connected {
			// Force connection to a neighbor (e.g. Up)
			if endY > 0 {
				grid[endY-1][endX].Type = models.Path
			}
		}
	} else {
		grid[endY][endX].Type = models.End
	}

	// 3. Add Questions to Path
	// [REMOVED] User requested removal of questions.
	/*
	   for y := 0; y < rows; y++ {
	       for x := 0; x < cols; x++ {
	           if grid[y][x].Type == models.Path {
	               // 5% chance (lower density for large maze)
	               if rand.Float32() < 0.05 {
	                   grid[y][x].HasQuestion = true
	                   grid[y][x].QuestionID = store.QuestionBank[rand.Intn(len(store.QuestionBank))].ID
	               }
	           }
	       }
	   }
	*/

	// No "Question Wall" or "Beyond" logic needed for Standard Maze.

	return models.Board{Grid: grid, Rows: rows, Cols: cols}
}

// Function to handle player movement.
func MovePlayer(game *models.GameState, direction string) (string, int, error) { // Changed return type to include QuestionID
	// 1. Calculate new coordinate based on direction.
	newPos := game.Player.CurrentPos
	switch direction {
	case "UP":
		newPos.Y--
	case "DOWN":
		newPos.Y++
	case "LEFT":
		newPos.X--
	case "RIGHT":
		newPos.X++
	}

	// 2. Check bounds.
	// If new coordinate is < 0 or >= size, return error (Invalid Move).
	if newPos.X < 0 || newPos.X >= game.Board.Cols || newPos.Y < 0 || newPos.Y >= game.Board.Rows {
		return "Invalid Move", -1, nil
	}

	// 3. Check cell type.
	cell := game.Board.Grid[newPos.Y][newPos.X]

	if cell.Type == models.Wall {
		return "Blocked", -1, nil
	}
	if cell.Type == models.End {
		game.Status = "WON"
		return "Win", -1, nil
	}

	// 4. If Path or Start:
	// Update Player.CurrentPos.
	// Check if this cell has a question? return "QuestionFound".
	game.Player.CurrentPos = newPos

	// Check for standard question on path
	if cell.HasQuestion {
		return "QuestionFound", cell.QuestionID, nil
	}

	return "Moved", -1, nil
}

// Function to check answer.
func CheckAnswer(game *models.GameState, questionID int, answer string) bool {
	// TODO: 1. Find the question from the store.
	if questionID < 0 {
		return false
	}
	// Simple lookup. Ideally use a map.
	var question *models.Question
	for _, q := range store.QuestionBank {
		if q.ID == questionID {
			question = &q
			break
		}
	}

	if question == nil {
		return false
	}

	// TODO: 2. Compare answer & 3. Logic.
	if question.CorrectAns == answer {
		game.Player.Score += 10
		return true
	}

	game.Player.Lives--
	if game.Player.Lives == 0 {
		game.Status = "LOST"
	}
	return false
}

// AnswerQuestion handles the logic for answering a question and returns the updated game state.
func AnswerQuestion(game *models.GameState, questionID int, answer string) (*AnswerResult, error) {
	// Validate QuestionID
	if questionID < 0 {
		return nil, fmt.Errorf("invalid question ID")
	}

	// Check the answer
	correct := CheckAnswer(game, questionID, answer)

	if correct {
		// Remove question from current pos if it was a path question
		p := game.Player.CurrentPos
		if game.Board.Grid[p.Y][p.X].HasQuestion && game.Board.Grid[p.Y][p.X].QuestionID == questionID {
			game.Board.Grid[p.Y][p.X].HasQuestion = false
			game.Board.Grid[p.Y][p.X].QuestionID = -1 // Reset QuestionID
		}
	}

	return &AnswerResult{
		Correct:   correct,
		GameState: game,
	}, nil
}
