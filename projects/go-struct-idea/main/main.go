package main

import (
	"fmt"
	structs "go-struct-idea/struct"
)

func main() {
	fmt.Println("This is a simple program implementation of structs in Go. For such, please choose between a car or a truck to start. ")
	fmt.Println("Take notice that the choice must be exactly 'car' or 'truck', or the program will turn to car by default.")

	var chc string
	fmt.Scan(&chc)

	if chc == "truck" {
		truck := structs.NewDefaultTruck()
		fmt.Println(truck)
	} else {
		car := structs.NewDefaultCar()
		fmt.Println(car)
	}
}
