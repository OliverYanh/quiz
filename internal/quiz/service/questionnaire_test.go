package service

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/oliveryanh/quiz/internal/quiz/model"
)

func setUp(testName string) func() {
	fmt.Printf("\tsetUp fixture for %s\n", testName)
	return func() {
		fmt.Printf("\ttearDown fixture for %s\n", testName)
	}
}

func TestFuncl(t *testing.T) {
	t.Cleanup(setUp(t.Name()))
	fmt.Printf("\tExecute test:%s\n", t.Name())
}

func pkgSetUp(pkgName string) func() {
	fmt.Printf("package SetUp fixture for %s\n", pkgName)
	return func() {
		fmt.Printf("package TearDown fixture for %s\n", pkgName)
	}
}
func TestMain(m *testing.M) {
	defer pkgSetUp("package test")
	m.Run()
}
func TestQuestionnaireService_LoadQuizData(t *testing.T) {
	// 准备测试数据
	quizData := &model.Quiz{
		ID: "123",
		Questions: []model.Question{
			{
				ID:            "q1",
				Text:          "What is the capital of France?",
				Answers:       []model.Answer{{ID: "a1", Text: "Paris"}, {ID: "a2", Text: "London"}, {ID: "a3", Text: "Berlin"}, {ID: "a4", Text: "Rome"}},
				CorrectAnswer: "a1",
			},
			// 添加更多的问题...
		},
	}
	// 将测试数据写入临时文件
	tempFile, err := os.CreateTemp("", "quiz_data_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name()) // 删除临时文件
	jsonEncoder := json.NewEncoder(tempFile)
	if err := jsonEncoder.Encode(quizData); err != nil {
		t.Fatal(err)
	}

	// 创建 QuestionnaireService 实例
	qs, _ := NewQuestionnaireService()

	// 加载测试数据
	err = qs.LoadQuizData(tempFile.Name())
	if err != nil {
		t.Errorf("LoadQuizData() returned unexpected error: %v", err)
	}

	// 检查加载的数据是否正确
	if len(qs.GetQuiz().Questions) != len(quizData.Questions) {
		t.Errorf("LoadQuizData() did not load the correct number of questions")
	}
}

func TestQuestionnaireService_SubmitUserAnswers(t *testing.T) {
	// 创建 QuestionnaireService 实例
	qs, _ := NewQuestionnaireService()

	// 设置测试数据
	quizData := &model.Quiz{
		ID: "123",
		Questions: []model.Question{
			{
				ID:            "q1",
				Text:          "What is the capital of France?",
				Answers:       []model.Answer{{ID: "a1", Text: "Paris"}, {ID: "a2", Text: "London"}, {ID: "a3", Text: "Berlin"}, {ID: "a4", Text: "Rome"}},
				CorrectAnswer: "a1",
			},
			// 添加更多的问题...
		},
	}
	qs.LoadQuizDataFromMemory(quizData) // 使用内存中的测试数据

	// 提交用户答案
	userAnswers := map[string]string{
		"q1": "a1", // 正确答案
		// 添加更多的用户答案...
	}
	score, _ := qs.SubmitUserAnswers(userAnswers)

	// 检查得分是否正确
	expectedScore := int32(1) // 此处假设用户只答对了一道题
	if score != expectedScore {
		t.Errorf("SubmitUserAnswers() returned incorrect score: got %d, want %d", score, expectedScore)
	}

	// 检查用户答案和得分是否正确存储
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
