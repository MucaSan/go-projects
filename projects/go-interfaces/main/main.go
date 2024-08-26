package main

import (
	inter "go-interfaces/inter"
)

func main() {
	car := inter.Car{Name: "car", NumTires: 4, Weight: 500.00}

	truck := inter.Truck{Name: "truck", NumTires: 4, Weight: 3000.00}

	motorcycle := inter.Motorcycle{Name: "motorcycle", NumTires: 2, Weight: 100.00}
}
