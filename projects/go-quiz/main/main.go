package main

import (
	csv "encoding/csv"
	qz "go-quiz/quiz"
	"strings"
)

func main() {
	s := qz.ReadCsvFile("../quiz/problems.csv")

	r := csv.NewReader(strings.NewReader(s))

	qz.PrintQuiz(r)
}
