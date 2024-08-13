package main

import (
	"fmt"
	"go-simple-arithmetic/aux_files"
)

func main() {
	fmt.Println("----------------------------------------------------------INSTRUCTIONS---------------------------------------------------------------------------------")
	fmt.Println("Hello, welcome to this simple calculator program!")
	fmt.Println("We won't have more than one try, as this simple project only takes into consideration really simple aspects of the language and it's purely educational")
	fmt.Println("If you don't use a valid operator, it will by default just multiple the given numbers.")
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------")
	var x, y float64
	var z string

	fmt.Println("Insert the first number: ")
	fmt.Scan(&x)

	fmt.Println("Insert the second number: ")
	fmt.Scan(&y)

	fmt.Println("Select the operation to be done('+', '-', '/' or '*'): ")
	fmt.Scan(&z)
	aux_files.ReturnOption(z, x, y)
}
