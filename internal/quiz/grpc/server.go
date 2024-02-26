package grpc

import (
	"context"

	"github.com/oliveryanh/quiz/internal/quiz/generated"
	"github.com/oliveryanh/quiz/internal/quiz/model"
	"github.com/oliveryanh/quiz/internal/quiz/service"
)

// QuizServer is a gRPC server for handling quiz-related requests.
type QuizServer struct {
	generated.UnimplementedQuizServiceServer
	service *service.QuestionnaireService
}

// NewQuizServer creates a new instance of QuizServer.
func NewQuizServer() *QuizServer {
	// Initialize the questionnaire service.
	svc, _ := service.NewQuestionnaireService()
	return &QuizServer{
		service: svc,
	}
}

// GetQuiz retrieves the quiz questions.
func (s *QuizServer) GetQuiz(ctx context.Context, req *generated.GetQuizRequest) (*generated.GetQuizResponse, error) {
	// Get the quiz from the service.
	quiz := s.service.GetQuiz()
	return &generated.GetQuizResponse{Quiz: convertToGRPCQuiz(quiz)}, nil
}

// SubmitQuiz submits the user's answers to the quiz.
func (s *QuizServer) SubmitQuiz(ctx context.Context, req *generated.SubmitQuizRequest) (*generated.SubmitQuizResponse, error) {
	// Convert gRPC question answers to a map.
	questionAnswers := convertToQuestionAnswers(req.QuestionAnswers)
	// Submit user answers and get the score and result.
	resp, result := s.service.SubmitUserAnswers(questionAnswers)

	// Return the response.
	return &generated.SubmitQuizResponse{Score: resp, Result: result}, nil
}

// convertToGRPCQuiz converts the internal quiz model to a gRPC-compatible format.
func convertToGRPCQuiz(quiz *model.Quiz) *generated.Quiz {
	var questions []*generated.Question
	for _, q := range quiz.Questions {
		var answers []*generated.Answer
		for _, a := range q.Answers {
			answers = append(answers, &generated.Answer{
				Id:   a.ID,
				Text: a.Text,
			})
		}
		questions = append(questions, &generated.Question{
			Id:            q.ID,
			Text:          q.Text,
			Answers:       answers,
			CorrectAnswer: q.CorrectAnswer,
		})
	}
	return &generated.Quiz{
		Questions: questions,
	}
}

// convertToQuestionAnswers converts gRPC question answers to a map.
func convertToQuestionAnswers(grpcQuestionAnswers []*generated.QuestionAnswer) map[string]string {
	questionAnswers := make(map[string]string)

	for _, grpcQA := range grpcQuestionAnswers {
		questionAnswers[grpcQA.QuestionId] = grpcQA.AnswerId
	}
	return questionAnswers
}
