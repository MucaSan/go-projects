package quiz

import (
	csv "encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
		fmt.Scanf("%s", &input)
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
	fmt.Println("Leaving it blank, will result in the program using the default CSV file. ")
	os.Exit(0)
}
