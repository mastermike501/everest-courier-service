package fleet

type Fleet struct {
	numVehicles        int
	maxSpeed           int
	maxCarriableWeight float64
}

func NewFleet(vehicles, speed int, weight float64) *Fleet {
	return &Fleet{
		numVehicles:        vehicles,
		maxSpeed:           speed,
		maxCarriableWeight: weight,
	}
}

func (f *Fleet) GetNumVehicles() int {
	return f.numVehicles
}

func (f *Fleet) GetMaxSpeed() int {
	return f.maxSpeed
}

func (f *Fleet) GetMaxCarriableWeight() float64 {
	return f.maxCarriableWeight
}
