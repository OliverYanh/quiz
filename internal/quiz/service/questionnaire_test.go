package service

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/oliveryanh/quiz/internal/quiz/model"
)

func TestQuestionnaireService_LoadQuizData(t *testing.T) {
	// Prepare test data
	quizData := &model.Quiz{
		ID: "123",
		Questions: []model.Question{
			{
				ID:            "q1",
				Text:          "What is the capital of France?",
				Answers:       []model.Answer{{ID: "a1", Text: "Paris"}, {ID: "a2", Text: "London"}, {ID: "a3", Text: "Berlin"}, {ID: "a4", Text: "Rome"}},
				CorrectAnswer: "a1",
			},
		},
	}
	// Write test data to temp file
	tempFile, err := os.CreateTemp("", "quiz_data_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name()) // Delete temp file
	jsonEncoder := json.NewEncoder(tempFile)
	if err := jsonEncoder.Encode(quizData); err != nil {
		t.Fatal(err)
	}

	// Create QuestionnaireService instance
	qs, _ := NewQuestionnaireService()

	// Load test data
	err = qs.LoadQuizData(tempFile.Name())
	if err != nil {
		t.Errorf("LoadQuizData() returned unexpected error: %v", err)
	}

	// Check if the loaded data is correct.
	if len(qs.GetQuiz().Questions) != len(quizData.Questions) {
		t.Errorf("LoadQuizData() did not load the correct number of questions")
	}
}

func TestQuestionnaireService_SubmitUserAnswers(t *testing.T) {
	// Create an instance of the QuestionnaireService.
	qs, _ := NewQuestionnaireService()

	// Set up test data.
	quizData := &model.Quiz{
		ID: "123",
		Questions: []model.Question{
			{
				ID:            "q1",
				Text:          "What is the capital of France?",
				Answers:       []model.Answer{{ID: "a1", Text: "Paris"}, {ID: "a2", Text: "London"}, {ID: "a3", Text: "Berlin"}, {ID: "a4", Text: "Rome"}},
				CorrectAnswer: "a1",
			},
		},
	}
	qs.LoadQuizDataFromMemory(quizData) // Use test data in memory.

	// Submit user answers.
	userAnswers := map[string]string{
		"q1": "a1", // Correct answers
	}
	score, _ := qs.SubmitUserAnswers(userAnswers)

	// Check if the score is correct.
	expectedScore := int32(1) // Here it is assumed that the user only answered one question correctly.
	if score != expectedScore {
		t.Errorf("SubmitUserAnswers() returned incorrect score: got %d, want %d", score, expectedScore)
	}

	// Check if user answers and scores are correctly stored.
	quiz := qs.GetQuiz()
	if quiz.Score != 1 {
		t.Errorf("SubmitUserAnswers() did not store correct score: got %d, want %d", quiz.Score, expectedScore)
	}
	for questionID, expectedAnswer := range userAnswers {
		if quiz.UserAnswers[questionID] != expectedAnswer {
			t.Errorf("SubmitUserAnswers() did not store correct user answer for question %s: got %s, want %s", questionID, quiz.UserAnswers[questionID], expectedAnswer)
		}
	}
}

func TestNewQuestionnaireService(t *testing.T) {

	qs, err := NewQuestionnaireService()

	if err != nil {
		t.Fatal("Unexpected error creating QuestionnaireService")
	}

	if qs.quiz == nil {
		t.Fatal("Quiz is nil after creating QuestionnaireService")
	}

	// Check the default number of loaded questions.
	if len(qs.quiz.Questions) != 3 {
		t.Errorf("Expected 3 default questions, got %d", len(qs.quiz.Questions))
	}
}

func TestLoadQuizData(t *testing.T) {

	qs, _ := NewQuestionnaireService()

	// Prepare test question data.
	quizData := &model.Quiz{
		Questions: []model.Question{
			{ID: "q1"},
			{ID: "q2"},
		},
	}

	// Test loading.
	qs.LoadQuizDataFromMemory(quizData)

	// Check the number of questions.
	if len(qs.GetQuiz().Questions) != 2 {
		t.Error("Failed to load expected number of questions")
	}
}
