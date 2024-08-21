package main

import (
	"fmt"
	qz "go-quiz/quiz"
	"log"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("This program is a QUIZ, in which it measures how answers have you got right and how many wrong. ")
		fmt.Println("The csv file from where the questions are taken is located at the quiz directory. ")
		s, err := qz.ReadCsvFile("../quiz/problems.csv")
		if err != nil {
			fmt.Println(`The default CSV file no longer exists. Please create another one with the name "problems.csv" on the quiz directory.`)
			log.Fatal()
		}
		r := qz.LoadReaderStruct(s)
		qz.PrintQuiz(r)
	} else if len(os.Args) == 2 && (os.Args[1] == "help" || os.Args[1] == "-h") {
		qz.DisplayHelpInfo()
	} else if len(os.Args) == 3 && (os.Args[1] == "--url" || os.Args[1] == "-u") {
		s, err := qz.ReadCsvFile(os.Args[2])
		if err != nil {
			log.Fatal()
		}
		fmt.Println("This program is a QUIZ, in which it measures how answers have you got right and how many wrong. ")
		fmt.Println("The csv file from where the questions are taken is located at the quiz directory. ")
		r := qz.LoadReaderStruct(s)
		qz.PrintQuiz(r)
	} else {
		qz.DisplayHelpInfo()
	}

}
