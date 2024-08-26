// package responsible for implemting interfaces which will be used on the main archive!

package inter

import (
	"fmt"
	"reflect"
)

type Vehicle interface {
	Drive()
	GetNumTires() int
	GetName() string
	GetWeight() float64
}

type Car struct {
	Name     string
	NumTires int
	Weight   float64
}

type Motorcycle struct {
	Name     string
	NumTires int
	Weight   float64
}

type Truck struct {
	Name     string
	NumTires int
	Weight   float64
}

func (t Truck) GetNumTires() int {
	return t.NumTires
}
func (t Truck) GetWeight() float64 {
	return t.Weight
}
func (t Truck) Drive() {
	fmt.Println("Driving the almighty truck!")
}
func (t Truck) GetName() string {
	return t.Name
}
func (c Car) Drive() {
	fmt.Println("Driving the car!")
}
func (c Car) GetName() string {
	return c.Name
}
func (c Car) GetWeight() float64 {
	return c.Weight
}
func (c Car) GetNumTires() int {
	return c.NumTires
}
func (m Motorcycle) GetNumTires() int {
	return m.NumTires
}
func (m Motorcycle) Drive() {
	fmt.Println("Driving the motorcycle!")
}
func (m Motorcycle) GetName() string {
	return m.Name
}
func (m Motorcycle) GetWeight() float64 {
	return m.Weight
}
func showInfo(v Vehicle) {
	tp := reflect.TypeOf(v).String()
	return ("type: " + tp + "\n" +
		"name: .")
}

func DescribeVehicle(v Vehicle) {
	switch v.(type) {
	case Motorcycle:
		fmt.Println()
	}
}
