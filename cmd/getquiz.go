// cmd/getquiz.go

package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/oliveryanh/quiz/internal/quiz/generated"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getQuizCmd)
}

// getQuizCmd represents the getquiz command.
var getQuizCmd = &cobra.Command{
	Use:     "getquiz",
	Aliases: []string{"gq"},
	Short:   "Get quiz questions",
	Run: func(cmd *cobra.Command, args []string) {
		getQuiz()
	},
}

// getQuiz retrieves and displays quiz questions.
func getQuiz() {
	conn, err := dialServer()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := generated.NewQuizServiceClient(conn)

	resp, err := client.GetQuiz(context.Background(), &generated.GetQuizRequest{})
	if err != nil {
		log.Fatalf("Failed to get quiz: %v", err)
	}

	fmt.Println("Quiz questions:")
	for i, question := range resp.Quiz.Questions {
		fmt.Printf("%d. Question: %s\n", i+1, question.Text)
		fmt.Println("Answers:")
		for _, answer := range question.Answers {
			fmt.Printf("   %s. %s\n", answer.GetId(), answer.Text)
		}
	}
}
