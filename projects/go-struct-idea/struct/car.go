package structs

type Car struct {
	Vehicle
}

func NewDefaultCar() Car {
	return Car{
		Vehicle: Vehicle{
			Model:       "car",
			Width:       3,
			Height:      1.5,
			NumberTires: 4,
		},
	}
}
