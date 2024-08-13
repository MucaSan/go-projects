package aux_files

import (
	"fmt"
)

func ReturnOption(op string, x, y float64) {
	var res float64
	if op == "+" {
		res = sum(x, y)
	} else if op == "-" {
		res = minus(x, y)
	} else if op == "/" {
		res = divide(x, y)
	} else {
		res = multiple(x, y)
	}
	fmt.Printf("The result is: %f ", res)
}
