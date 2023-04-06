package fleet

import (
	"math"
	"testing"
)

func TestNewVehicle(t *testing.T) {
	name := "Car"
	returnTime := 2.5
	v := NewVehicle(name, returnTime)
	if v.name != name {
		t.Errorf("NewVehicle(%v, %v) failed. Expected name=%v, but got name=%v", name, returnTime, name, v.name)
	}
	if v.returnTime != returnTime {
		t.Errorf("NewVehicle(%v, %v) failed. Expected returnTime=%v, but got returnTime=%v", name, returnTime, returnTime, v.returnTime)
	}
}

func TestVehicleSetReturnTime(t *testing.T) {
	name := "Car"
	returnTime := 2.5
	newReturnTime := 4.5
	v := NewVehicle(name, returnTime)
	v.SetReturnTime(newReturnTime)
	if v.returnTime != newReturnTime {
		t.Errorf("Vehicle.SetReturnTime(%v) failed. Expected returnTime=%v, but got returnTime=%v", newReturnTime, newReturnTime, v.returnTime)
	}
}

func TestVehicleGetReturnTime(t *testing.T) {
	name := "Car"
	returnTime := 2.5
	v := NewVehicle(name, returnTime)
	if v.GetReturnTime() != returnTime {
		t.Errorf("Vehicle.GetReturnTime() failed. Expected returnTime=%v, but got returnTime=%v", returnTime, v.GetReturnTime())
	}
}

func TestVehicleUpdateReturnTime(t *testing.T) {
	name := "Car"
	returnTime := 2.5
	curTime := 1.5
	v := NewVehicle(name, returnTime)
	v.UpdateReturnTime(curTime)
	expectedReturnTime := returnTime - curTime
	if math.Abs(v.returnTime-expectedReturnTime) > 0.0001 {
		t.Errorf("Vehicle.UpdateReturnTime(%v) failed. Expected returnTime=%v, but got returnTime=%v", curTime, expectedReturnTime, v.returnTime)
	}

	// Test negative return time
	v.SetReturnTime(0.5)
	v.UpdateReturnTime(1.0)
	if v.returnTime != 0 {
		t.Errorf("Vehicle.UpdateReturnTime(%v) failed. Expected returnTime=0, but got returnTime=%v", curTime, v.returnTime)
	}
}
