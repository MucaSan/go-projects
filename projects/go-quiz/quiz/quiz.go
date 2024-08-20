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

func ReadCsvFile(url string) string {
	b, err := os.ReadFile(url)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)

}
