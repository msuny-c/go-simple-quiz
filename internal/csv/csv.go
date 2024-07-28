package csv

import (
	"quiz/internal/model"
	"encoding/csv"
	"os"
)

func ReadQuizzes(file *os.File) (quizzes []model.Quiz) {
	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.FieldsPerRecord = 2
	for {
		record, end := reader.Read()
		if end != nil {
			break
		}
		quizzes = append(quizzes, model.Quiz{
			Question: record[0], 
			Answer: record[1],
		})
	}
	return
}
