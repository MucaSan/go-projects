package struct

type Truck struct {
	Vehicle
}

func NewDefaultTruck() Truck{
	return Truck{
		Model:"truck"
		Width:12
		Height: 2.2
		NumberTires: 4
	}
}

func NewCustomTruck(model string, width, height float64, numberTires int) Truck{
	return Truck{
		Model: model
		Width: width
		Height: height
		NumberTires: numberTires
	}
}