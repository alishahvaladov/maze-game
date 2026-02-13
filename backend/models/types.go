package models

// Define your data structures here. Start simple.

// Position represents x,y coordinates on the grid.
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// CellType represents what is in a cell (Wall, Path, Start, End).
type CellType string

const (
	Path   CellType = "PATH"
	Wall   CellType = "WALL"
	Start  CellType = "START"
	End    CellType = "EXIT"
	Beyond CellType = "BEYOND"
)

// Cell represents a single block in the grid.
type Cell struct {
	Type        CellType `json:"type"`
	Position    Position `json:"position"`
	HasQuestion bool     `json:"has_question"`
	// If it has a question, you might store the QuestionID here.
	QuestionID int `json:"question_id,omitempty"`
	// Is it a special "Question Wall"?
	IsQuestionWall bool `json:"is_question_wall"`
}

// Board represents the grid.
type Board struct {
	Rows int      `json:"rows"`
	Cols int      `json:"cols"`
	Grid [][]Cell `json:"grid"`
}

// Player represents the user's state.
type Player struct {
	CurrentPos Position `json:"current_pos"`
	Lives      int      `json:"lives"`
	Score      int      `json:"score"`
}

// GameState represents the entire state of a single game session.
type GameState struct {
	ID     string `json:"id"`
	Board  Board  `json:"board"`
	Player Player `json:"player"`
	Status string `json:"status"` // "ACTIVE", "WON", "LOST"
}

// Question represents a quiz question.
type Question struct {
	ID         int      `json:"id"`
	Text       string   `json:"text"`
	Options    []string `json:"options"`
	CorrectAns string   `json:"correct_ans"` // Should probably not send this to frontend immediately!
	Difficulty string   `json:"difficulty"`
}
