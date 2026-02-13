package store

import (
	"encoding/json"
	"fmt"
	"maze-game/models"
	"os"
)

// Global variable to hold loaded questions.
var QuestionBank []models.Question

// LoadQuestions reads the questions.json file and populates the QuestionBank.
func LoadQuestions(filePath string) error {
	// TODO: 1. Open the file.
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// TODO: 2. Decode JSON.
	err = json.Unmarshal(file, &QuestionBank)
	if err != nil {
		return err
	}

	// TODO: 3. Handle errors.
	fmt.Println("Questions loaded successfully", QuestionBank)

	return nil
}
