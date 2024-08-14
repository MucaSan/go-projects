package struct

type Car struct {
	Vehicle
	DistanceTraveled float64
}

func NewDefaultCar() Car{
	return Car{
		Model:"car"
		Width:3
		Height: 1.5
		NumberTires: 4
	}
}

func NewCustomCar(model string, width, height float64, numberTires int) Car{
	return Car{
		Model: model
		Width: width
		Height: height
		NumberTires: numberTires
	}
}