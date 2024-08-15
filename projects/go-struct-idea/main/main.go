package main

import (
	"fmt"
	structs "go-struct-idea/struct"
)

func main() {
	fmt.Println("it's working")
	car := structs.NewDefaultCar()
	fmt.Println(car)
}
