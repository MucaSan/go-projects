package main

import "fmt"

func main() {
	// manually set the value, based on the number inserted.
	var number1 int = 32
	// automatically instantiates the type of the variable, based on the value the programmer has inputed.
	// which basically means is that the ':=' operator instances the type, based on the value inputed by the programmer
	number2 := 32
	//  const cannot hold hte ":=" operator, as the
	const number3 int = 4
	// Interpolate strings, and return something.
	// You can use %d for ints, %f for float values, %s for string and  %v for any type  (not recommended if you know the type of the variable	)
	s := fmt.Sprintf("Hello, test! %s ", "conc")

	fmt.Println(number1)
	fmt.Println(number2)
	fmt.Println(number3)
	fmt.Printf(s)

}
