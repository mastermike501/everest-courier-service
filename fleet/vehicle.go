package fleet

type Vehicle struct {
	Name       string
	ReturnTime float64
}

func (v *Vehicle) UpdateReturnTime(curTime float64) {
	// if the calculated return time is negative, it means the vehicle
	// has already reached back to depot. Assign ReturnTime to zero
	returnTime := v.ReturnTime - curTime
	if returnTime < 0 {
		v.ReturnTime = 0
		return
	}

	// else, the returnTime would be the time remaining
	v.ReturnTime = returnTime
}
