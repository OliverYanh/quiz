package model

import (
	"encoding/json"
	"os"
)

// Answer represents a possible answer to a question
type Answer struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// Question represents a quiz question
type Question struct {
	ID            string   `json:"id"`
	Text          string   `json:"text"`
	Answers       []Answer `json:"answers"`
	CorrectAnswer string   `json:"correctAnswer"`
}

// Quiz represents a quiz taken by a user.
type Quiz struct {
	ID          string
	Questions   []Question        `json:"questions"`
	UserAnswers map[string]string // Maps question ID to the ID of the answer selected by the user
	Score       int32
}

// LoadData loads quiz data from a JSON file
func LoadData(filename string) (*Quiz, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var quizData struct {
		Questions []Question `json:"questions"`
		Quiz      struct {
			Questions []string `json:"questions"`
		} `json:"quiz"`
	}

	if err := json.Unmarshal(data, &quizData); err != nil {
		return nil, err
	}

	quiz := NewQuiz()
	for _, qData := range quizData.Questions {
		quiz.AddQuestion(qData)
	}

	return quiz, nil
}

// AddQuestion adds a new question to the quiz
func (q *Quiz) AddQuestion(question Question) {
	q.Questions = append(q.Questions, question)
}

// NewQuiz creates a new instance of Quiz
func NewQuiz() *Quiz {
	return &Quiz{}
}
