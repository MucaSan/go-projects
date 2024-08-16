package structs

type Truck struct {
	Vehicle
	DistanceTraveled float64
}

func NewDefaultTruck() Truck {
	return Truck{
		Vehicle: Vehicle{
			Model:       "truck",
			Width:       12.2,
			Height:      2.2,
			NumberTires: 4,
		},
		DistanceTraveled: 0.0,
	}
}

func NewCustomTruck(model string, width, height, distanceTraveled float64, numberTires int) Truck {
	return Truck{
		Vehicle: Vehicle{
			Model:       model,
			Width:       width,
			Height:      height,
			NumberTires: numberTires,
		},
		DistanceTraveled: distanceTraveled,
	}
}
func (truck Truck) SetModel(model string) {
	truck.Vehicle.Model = model
}
func (truck Truck) Model() string {
	return truck.Vehicle.Model
}
func (truck Truck) Width() float64 {
	return truck.Vehicle.Width
}

func (truck Truck) SetWidth(width float64) {
	truck.Vehicle.Width = width
}

func (truck Truck) SetHeight(height float64) {
	truck.Vehicle.Height = height
}

func (truck Truck) Height() float64 {
	return truck.Vehicle.Height
}

func (truck Truck) SetNumberTires(numberTires int) {
	truck.Vehicle.NumberTires = numberTires
}

func (truck Truck) NumberTires() int {
	return truck.Vehicle.NumberTires
}

func (truck Truck) GetDistanceTraveled() float64 {
	return truck.DistanceTraveled
}

func (truck Truck) SetDistanceTraveled(distanceTraveled float64) {
	truck.DistanceTraveled = distanceTraveled
}
