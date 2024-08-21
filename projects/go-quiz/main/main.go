package main

import (
	qz "go-quiz/quiz"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		qz.DisplayQuizIntroduction()
		s, _ := qz.ReadCsvFile("../quiz/problems.csv")
		r := qz.LoadReaderStruct(s)
		qz.PrintQuiz(r)
	} else if len(os.Args) == 2 && (os.Args[1] == "help" || os.Args[1] == "-h") {
		qz.DisplayHelpInfo()
	} else if len(os.Args) == 3 && (os.Args[1] == "--url" || os.Args[1] == "-u") {
		s, _ := qz.ReadCsvFile(os.Args[2])
		qz.DisplayQuizIntroduction()
		r := qz.LoadReaderStruct(s)
		qz.PrintQuiz(r)
	} else {
		qz.DisplayHelpInfo()
	}
}
