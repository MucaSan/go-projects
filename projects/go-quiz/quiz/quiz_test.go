package quiz

import (
	"fmt"
	"io"
	"testing"
)

func TestReadCsv(t *testing.T) {
	test1, _ := ReadCsvFile("../quiz/test1.csv")
	test2, _ := ReadCsvFile("../quiz/test2.csv")
	test3, _ := ReadCsvFile("../quiz/test3.csv")
	if test1 != `compare_String?, "very big string"` {
		t.Errorf("ERROR AT TEST 1 - got %s, expected %s", test1, `compare_String?, "very big string"`)
	}
	if test2 != `""big test?"", 1` {
		t.Errorf("ERROR AT TEST 2 - got %s, expected %s", test2, `""big test?"", 1`)
	}
	if test3 != `"what's up", "1"` {
		t.Errorf("ERROR AT TEST 3 - got %s, expected %s", test2, `"what's up", "1"`)
	}
}

func TestLoadReaderStruct(t *testing.T) {
	fmt.Println("Error has been here")
	test1, _ := ReadCsvFile("../quiz/test1.csv")
	test2, _ := ReadCsvFile("../quiz/test2.csv")
	test3, _ := ReadCsvFile("../quiz/test3.csv")
	fmt.Println("Error has been here2")
	r1 := LoadReaderStruct(test1)
	r2 := LoadReaderStruct(test2)
	r3 := LoadReaderStruct(test3)
	fmt.Println("Error has been here3")
	for {

		rec1, err1 := r1.Read()
		rec2, err2 := r2.Read()
		rec3, err3 := r3.Read()
		fmt.Println("Error has been here4")
		if err1 == io.EOF || err2 == io.EOF || err3 == io.EOF {
			break
		}
		fmt.Println("Error has been here5")
		if err1 != nil || err2 != nil || err3 != nil {
			t.Errorf("One of the errors from read has retrieved an error.")
		}
		fmt.Println("Error has been here6")
		if rec1[0] != "compare_String?" || rec1[1] != `"very big string"` {
			t.Errorf("First element: %s and Second element: %s. Expected %s and %s, respectively.", rec1[0], rec1[1], "compare_String?", `"very big string"`)
		}
		fmt.Println("Error has been here7")
		if rec2[0] != `""big test?""` || rec2[1] != "1" {
			t.Errorf("First element: %s and Second element: %s. Expected %s and %s, respectively.", rec2[0], rec2[1], `""big test?""`, "1")
		}
		fmt.Println("Error has been here8")
		if rec3[0] != `"what's up"` || rec3[1] != `"1"` {
			t.Errorf("First element: %s and Second element: %s. Expected %s and %s, respectively.", rec3[0], rec3[1], `"what's up"`, `"1"`)
		}
		fmt.Println("Error has been here9")
	}

}
