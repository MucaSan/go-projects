package main

import (
	"fmt"
)

type Test struct {
	name string
}

func (t *Test) SetName(name string) {
	t.name = name
}

func (t Test) GetName() string {
	return t.name
}

func main() {
	fmt.Println("It's working")
	t := new(Test)

	t.SetName("test")

	fmt.Println(t.GetName())
}
