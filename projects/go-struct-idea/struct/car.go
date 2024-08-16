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
func NewCustomCar(model string, width, height float64, numberTires int) Car {
	return Car{
		Vehicle: Vehicle{
			Model:       model,
			Width:       width,
			Height:      height,
			NumberTires: numberTires,
		},
	}
}

func (car Car) SetModel(model string) {
	car.Vehicle.Model = model
}
func (car Car) Model() string {
	return car.Vehicle.Model
}
func (car Car) Width() float64 {
	return car.Vehicle.Width
}

func (car Car) SetWidth(width float64) {
	car.Vehicle.Width = width
}

func (car Car) SetHeight(height float64) {
	car.Vehicle.Height = height
}

func (car Car) Height() float64 {
	return car.Vehicle.Height
}

func (car Car) SetNumberTires(numberTires int) {
	car.Vehicle.NumberTires = numberTires
}

func (car Car) NumberTires() int {
	return car.Vehicle.NumberTires
}
