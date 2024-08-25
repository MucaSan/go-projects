package quiz

import (
	"bufio"
	csv "encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	QuestionTest1 string = "compare_String?"
	QuestionTest2 string = "big test? 'beautiful'"
	QuestionTest3 string = "what's up"
	QuestionTest4 string = "123"
	AnswerTest1   string = "3"
	AnswerTest2   string = "1"
	AnswerTest3   string = `"abcde"`
	AnswerTest4   string = "'123'"
)

func showResults(totalAnswers, rightAnswers int) {
	fmt.Println("Congratulations, you've finished the quiz!")
	fmt.Printf("You've got %d answers right, out of %d !\n", rightAnswers, totalAnswers)
}

func PrintQuiz(r *csv.Reader) {
	var input string
	cnt := 0
	cntCorrect := 0
	for {

		rec, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d -  %s ? \n", cnt+1, rec[0])
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		if input == rec[1] {
			cntCorrect += 1
		}
		cnt += 1
	}
	showResults(cnt, cntCorrect)
}

func LoadReaderStruct(s string) *csv.Reader {
	return csv.NewReader(strings.NewReader(s))
}

func ReadCsvFile(url string) (string, error) {
	b, err := os.ReadFile(url)
	if err != nil {
		log.Fatal(err)
	}
	return string(b), err

}

func DisplayHelpInfo() {
	fmt.Println("Usage: go-quiz <command> ")
	fmt.Println("\n Options [-u, or --url] and the URL text from which you can set CSV file for the quiz.")
	fmt.Println("The CSV syntax file must be the same as the one in the quiz directory.")
	fmt.Println("Leaving it blank, will result in the program using the default CSV file. \n \n ")
	fmt.Println("Format CSV syntax should follow STRICTLY the given CSV tests formats. ")
	fmt.Println("The answers or questions should be done under a certain pattern, such as regarded in the text below.")
	fmt.Println(`1. "" Usage`)
	fmt.Println(`1.1 Whenever using "" with ', the text should be explicted as: " 'test' " and NOT 'test' by itself. `)
	fmt.Println(`1.2 Whenever using double quotes for enclosing fields, in the CSV text, the syntax should be: " ""test"" " and`)
	fmt.Println(`"NOT " "test" ", for example.`)
	fmt.Println("2. Program usage. \n \n ")
	fmt.Println("Be cautious about the CSV syntax file and take try testing your own quiz, before passing it to someone else. ")
	os.Exit(0)
}

func DisplayQuizIntroduction() {
	fmt.Println("This program is a QUIZ, in which it measures how answers have you got right and how many wrong. ")
	fmt.Println("The csv file from where the questions are taken is located at the quiz directory. ")
}
