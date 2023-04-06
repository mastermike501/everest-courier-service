package fleet

type Vehicle struct {
	name       string
	returnTime float64
}

func NewVehicle(name string, returnTime float64) *Vehicle {
	return &Vehicle{
		name:       name,
		returnTime: returnTime,
	}
}

func (v *Vehicle) SetReturnTime(returnTime float64) {
	v.returnTime = returnTime
}

func (v *Vehicle) GetReturnTime() float64 {
	return v.returnTime
}

func (v *Vehicle) UpdateReturnTime(curTime float64) {
	// if the calculated return time is negative, it means the vehicle
	// has already reached back to depot. Assign ReturnTime to zero
	returnTime := v.returnTime - curTime
	if returnTime < 0 {
		v.returnTime = 0
		return
	}

	// else, the returnTime would be the time remaining
	v.returnTime = returnTime
}
