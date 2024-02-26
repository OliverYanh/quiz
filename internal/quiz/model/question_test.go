// question_test.go
package model

import (
	"reflect"
	"testing"
)

func TestLoadData(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     *Quiz
		wantErr  bool
	}{
		{
			name:     "valid data",
			filename: "testdata/quiz_data.json",
			want: &Quiz{
				Questions: []Question{
					{
						ID:            "q1",
						Text:          "What is the capital of France?",
						Answers:       []Answer{{ID: "a1", Text: "Paris"}, {ID: "a2", Text: "London"}, {ID: "a3", Text: "Berlin"}, {ID: "a4", Text: "Rome"}},
						CorrectAnswer: "a1",
					},
					{
						ID:            "q2",
						Text:          "What is 2 + 2?",
						Answers:       []Answer{{ID: "a5", Text: "3"}, {ID: "a6", Text: "4"}, {ID: "a7", Text: "5"}, {ID: "a8", Text: "6"}},
						CorrectAnswer: "a6",
					},
					{
						ID:            "q3",
						Text:          "What is the largest planet in our solar system?",
						Answers:       []Answer{{ID: "a9", Text: "Mars"}, {ID: "a10", Text: "Venus"}, {ID: "a11", Text: "Jupiter"}, {ID: "a12", Text: "Saturn"}},
						CorrectAnswer: "a11",
					},
				},
			},
			wantErr: false,
		},
		// Add more test cases for invalid data, missing file, etc. if needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadData(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadData() = %v, want %v", got, tt.want)
			}
		})
	}
}
