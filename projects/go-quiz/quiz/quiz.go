package quiz

import (
	csv "encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func PrintQuiz(r *csv.Reader) {
	for {
		rec, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(rec)
	}
}

func ReadCsvFile(url string) string {
	b, err := os.ReadFile("../quiz/problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)

}
