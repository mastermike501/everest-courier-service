package fleet

import (
	"testing"
)

func TestNewFleet(t *testing.T) {
	numVehicles := 5
	maxSpeed := 60
	maxCarriableWeight := 100.0

	f := NewFleet(numVehicles, maxSpeed, maxCarriableWeight)

	if f == nil {
		t.Error("Expected a non-nil Fleet object, but got nil")
	}

	if f.GetNumVehicles() != numVehicles {
		t.Errorf("Expected Fleet to have %d vehicles, but got %d", numVehicles, f.GetNumVehicles())
	}

	if f.GetMaxSpeed() != maxSpeed {
		t.Errorf("Expected Fleet to have max speed %d, but got %d", maxSpeed, f.GetMaxSpeed())
	}

	if f.GetMaxCarriableWeight() != maxCarriableWeight {
		t.Errorf("Expected Fleet to have max carriable weight %f, but got %f", maxCarriableWeight, f.GetMaxCarriableWeight())
	}
}

func TestFleetMethods(t *testing.T) {
	numVehicles := 5
	maxSpeed := 60
	maxCarriableWeight := 100.0

	f := NewFleet(numVehicles, maxSpeed, maxCarriableWeight)

	if f.GetNumVehicles() != numVehicles {
		t.Errorf("Expected Fleet to have %d vehicles, but got %d", numVehicles, f.GetNumVehicles())
	}

	if f.GetMaxSpeed() != maxSpeed {
		t.Errorf("Expected Fleet to have max speed %d, but got %d", maxSpeed, f.GetMaxSpeed())
	}

	if f.GetMaxCarriableWeight() != maxCarriableWeight {
		t.Errorf("Expected Fleet to have max carriable weight %f, but got %f", maxCarriableWeight, f.GetMaxCarriableWeight())
	}
}
