package service

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/oliveryanh/quiz/internal/quiz/model"
)

//go:embed data/quiz_data.json
var quizData []byte

// QuestionnaireService represents a service for managing questionnaires.
type QuestionnaireService struct {
	mu     sync.Mutex  // mu guards access to quiz and scores.
	quiz   *model.Quiz // quiz holds the current quiz data.
	scores []int32     // scores holds the scores of all users.
}

// NewQuestionnaireService creates a new instance of QuestionnaireService.
func NewQuestionnaireService() (*QuestionnaireService, error) {
	// Load quiz data from the embedded JSON file.
	var quiz model.Quiz
	if err := json.Unmarshal(quizData, &quiz); err != nil {
		return nil, fmt.Errorf("failed to parse embedded JSON: %v", err)
	}

	return &QuestionnaireService{
		quiz: &quiz,
	}, nil
}

// LoadQuizData loads quiz data from a JSON file and updates the current quiz data.
func (qs *QuestionnaireService) LoadQuizData(filename string) error {
	qs.mu.Lock()
	defer qs.mu.Unlock()

	// Resolve the absolute path to the data file.
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("failed to resolve absolute path: %v", err)
	}

	// Read JSON data from the file.
	data, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// Unmarshal JSON data into Quiz struct.
	err = json.Unmarshal(data, &qs.quiz)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	return nil
}

// GetQuiz returns the current quiz questions.
func (qs *QuestionnaireService) GetQuiz() *model.Quiz {
	qs.mu.Lock()
	defer qs.mu.Unlock()

	return qs.quiz
}

// SubmitUserAnswers submits user answers to the quiz and returns the score and comparison information.
func (qs *QuestionnaireService) SubmitUserAnswers(userAnswers map[string]string) (int32, string) {
	qs.mu.Lock()
	defer qs.mu.Unlock()

	// Calculate score.
	score := int32(0)
	for _, question := range qs.quiz.Questions {
		if userAnswer, ok := userAnswers[question.ID]; ok && userAnswer == question.CorrectAnswer {
			score++
		}
	}

	// Update quiz data with user answers and score.
	qs.quiz.UserAnswers = userAnswers
	qs.quiz.Score = score
	qs.scores = append(qs.scores, score)

	// Calculate the percentage of users that the current user performed better than.
	percentage := qs.calculateComparison(score)
	return score, fmt.Sprintf("You were better than %.2f%% of all quizzers", percentage)
}

// LoadQuizDataFromMemory loads quiz data from memory and updates the current quiz data.
func (qs *QuestionnaireService) LoadQuizDataFromMemory(quizData *model.Quiz) {
	qs.mu.Lock()
	defer qs.mu.Unlock()

	qs.quiz = quizData
}

// calculateComparison calculates the user's performance relative to other participants.
func (qs *QuestionnaireService) calculateComparison(userScore int32) float64 {
	numOthers := len(qs.scores)
	betterThan := 0

	for _, score := range qs.scores {
		if userScore >= score {
			betterThan++
		}
	}

	// Calculate the percentage of users the current user performed better than.
	betterThanPerc := float64(betterThan) / float64(numOthers) * 100

	return betterThanPerc
}
