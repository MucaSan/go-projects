package quiz

import (
	"io"
	"testing"
)

func TestReadCsv(t *testing.T) {
	test1, _ := ReadCsvFile("../quiz/test1.csv")
	test2, _ := ReadCsvFile("../quiz/test2.csv")
	test3, _ := ReadCsvFile("../quiz/test3.csv")
	test4, _ := ReadCsvFile("../quiz/test4.csv")
	if test1 != `"compare_String?",3` {
		t.Errorf("ERROR AT TEST 1 - got %s, expected %s", test1, `"compare_String?",3`)
	}
	if test2 != `"big test? 'beautiful'","1"` {
		t.Errorf("ERROR AT TEST 2 - got %s, expected %s", test2, `"big test? 'beautiful'","1"`)
	}
	if test3 != `"what's up","""abcde"""` {
		t.Errorf("ERROR AT TEST 3 - got %s, expected %s", test2, `"what's up","""abcde"""`)
	}
	if test4 != `123,"'123'"` {
		t.Errorf("ERROR AT TEST 3 - got %s, expected %s", test2, `123,"'123'"`)
	}
}

func TestLoadReaderStructCSV1(t *testing.T) {
	s, _ := ReadCsvFile("../quiz/test1.csv")
	r := LoadReaderStruct(s)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(rec) != 2 {
			t.Errorf(`Error, file reader has issues. File reader exited with %d size. \n 
			The string formated was {%s}, which caused the issue. `, len(rec), s)
		}
		if rec[0] != QuestionTest1 || rec[1] != AnswerTest1 {
			t.Errorf("First element: %s and Second element: %s. Expected %s and %s, respectively.", rec[0], rec[1], QuestionTest1, AnswerTest1)
		}
	}
}
func TestLoadReaderStructCSV2(t *testing.T) {
	s, _ := ReadCsvFile("../quiz/test2.csv")
	r := LoadReaderStruct(s)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(rec) != 2 {
			t.Errorf(`Error, file reader has issues. File reader exited with %d size. \n 
			The string formated was {%s}, which caused the issue. `, len(rec), s)
		}
		if rec[0] != QuestionTest2 || rec[1] != AnswerTest2 {
			t.Errorf("First element: %s and Second element: %s. Expected %s and %s, respectively.", rec[0], rec[1], QuestionTest2, AnswerTest2)
		}
	}
}

func TestLoadReaderStructCSV3(t *testing.T) {
	s, _ := ReadCsvFile("../quiz/test3.csv")
	r := LoadReaderStruct(s)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(rec) != 2 {
			t.Errorf(`Error, file reader has issues. File reader exited with %d size. \n 
			The string formated was {%s}, which caused the issue. `, len(rec), s)
		}
		if rec[0] != QuestionTest3 || rec[1] != AnswerTest3 {
			t.Errorf("First element: %s and Second element: %s. Expected %s and %s, respectively.", rec[0], rec[1], QuestionTest3, AnswerTest3)
		}
	}
}

func TestLoadReaderStructCSV4(t *testing.T) {
	s, _ := ReadCsvFile("../quiz/test4.csv")
	r := LoadReaderStruct(s)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(rec) != 2 {
			t.Errorf(`Error, file reader has issues. File reader exited with %d size. \n 
			The string formated was {%s}, which caused the issue. `, len(rec), s)
		}
		if rec[0] != QuestionTest4 || rec[1] != AnswerTest4 {
			t.Errorf("First element: %s and Second element: %s. Expected %s and %s, respectively.", rec[0], rec[1], QuestionTest4, AnswerTest4)
		}
	}
}
