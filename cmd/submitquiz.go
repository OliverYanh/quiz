package cmd

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/oliveryanh/quiz/internal/quiz/generated"
	"github.com/spf13/cobra"
)

// init initializes the submitquiz command and flags.
func init() {
	rootCmd.AddCommand(submitQuizCmd)
	submitQuizCmd.Flags().StringP("answers", "a", "", "Comma-separated list of answers (e.g., '1,2,3')")
	submitQuizCmd.MarkFlagRequired("answers")
}

var submitQuizCmd = &cobra.Command{
	Use:     "submitquiz [answers]",
	Aliases: []string{"sq"},
	Short:   "Submit quiz answers",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var answersStr string
		if len(args) == 1 {
			answersStr = args[0]
		} else {
			var err error
			answersStr, err = cmd.Flags().GetString("answers")
			if err != nil {
				fmt.Println("Failed to retrieve answers:", err)
				return
			}
		}

		submitQuiz(answersStr)
	},
}

// submitQuiz submits the quiz answers to the server.
func submitQuiz(answersStr string) {
	conn, err := dialServer()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := generated.NewQuizServiceClient(conn)

	answers := strings.Split(answersStr, ",")
	var questionAnswers []*generated.QuestionAnswer
	for i, ans := range answers {
		answerIdx, err := strconv.Atoi(ans)
		if err != nil {
			log.Fatalf("Invalid answer format: %v", err)
		}
		if answerIdx < 1 || answerIdx > 4 { // Assuming each question has 4 answers
			log.Fatalf("Invalid answer index for question %d: %d", i+1, answerIdx)
		}
		questionAnswers = append(questionAnswers, &generated.QuestionAnswer{
			QuestionId: strconv.Itoa(i + 1), // Assuming questions are indexed starting from 1
			AnswerId:   fmt.Sprintf("%d", answerIdx),
		})
	}

	submitQuizRequest := &generated.SubmitQuizRequest{
		QuestionAnswers: questionAnswers,
	}

	resp, err := client.SubmitQuiz(context.Background(), submitQuizRequest)
	if err != nil {
		log.Fatalf("Failed to submit quiz: %v", err)
	}

	fmt.Printf("Score: %d\n", resp.GetScore())
	fmt.Printf("Result: %s\n", resp.GetResult())
}
