package structs

type Truck struct {
	Vehicle
}

func NewDefaultTruck() Truck {
	return Truck{
		Vehicle: Vehicle{
			Model:       "truck",
			Width:       12.2,
			Height:      2.2,
			NumberTires: 4,
		},
	}
}

func NewCustomTruck(model string, width, height float64, numberTires int) Truck {
	return Truck{
		Vehicle: Vehicle{
			Model:       model,
			Width:       width,
			Height:      height,
			NumberTires: numberTires,
		},
	}
}
