package main

import (
	"fmt"
	qz "go-quiz/quiz"
)

func main() {
	fmt.Println("This program is a QUIZ, in which it measures how answers have you got right and how many wrong. ")
	fmt.Println("The csv file from where the questions are taken is located at the quiz directory. ")
	s := qz.ReadCsvFile("../quiz/problems.csv")

	r := qz.LoadReaderStruct(s)

	qz.PrintQuiz(r)
}
