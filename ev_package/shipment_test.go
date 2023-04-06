package ev_package

import (
	"testing"
)

func TestNewShipment(t *testing.T) {
	shipment := NewShipment()
	if shipment.oneWayDeliveryTime != 0.0 {
		t.Errorf("Expected oneWayDeliveryTime to be 0.0, but got %f", shipment.oneWayDeliveryTime)
	}
	if len(shipment.packages) != 0 {
		t.Errorf("Expected packages map to be empty, but got length %d", len(shipment.packages))
	}
}

func TestShipment_AddPackages(t *testing.T) {
	shipment := NewShipment()
	package1 := &Package{Name: "pkg1", Distance: 100.0, TimeToDest: 2.5}
	package2 := &Package{Name: "pkg2", Distance: 200.0, TimeToDest: 5.0}

	shipment.AddPackages([]*Package{package1, package2})

	if len(shipment.packages) != 2 {
		t.Errorf("Expected packages map to have length 2, but got length %d", len(shipment.packages))
	}

	expectedDeliveryTime := 5.0 // since package2 has a delivery time of 5.0
	if shipment.oneWayDeliveryTime != expectedDeliveryTime {
		t.Errorf("Expected oneWayDeliveryTime to be %f, but got %f", expectedDeliveryTime, shipment.oneWayDeliveryTime)
	}
}

func TestShipment_GetPackages(t *testing.T) {
	shipment := NewShipment()
	package1 := &Package{Name: "pkg1", Distance: 100.0, TimeToDest: 2.5}
	package2 := &Package{Name: "pkg2", Distance: 200.0, TimeToDest: 5.0}

	shipment.AddPackages([]*Package{package1, package2})

	packages := shipment.GetPackages()
	if len(packages) != 2 {
		t.Errorf("Expected packages map to have length 2, but got length %d", len(packages))
	}
	if packages["pkg1"] != package1 {
		t.Errorf("Expected packages[\"pkg1\"] to be package1, but got %v", packages["pkg1"])
	}
	if packages["pkg2"] != package2 {
		t.Errorf("Expected packages[\"pkg2\"] to be package2, but got %v", packages["pkg2"])
	}
}

func TestShipment_GetOneWayDeliveryTime(t *testing.T) {
	shipment := NewShipment()
	package1 := &Package{Name: "pkg1", Distance: 100.0, TimeToDest: 2.5}
	package2 := &Package{Name: "pkg2", Distance: 200.0, TimeToDest: 5.0}

	shipment.AddPackages([]*Package{package1, package2})

	expectedDeliveryTime := 5.0 // since package2 has a delivery time of 5.0
	if shipment.GetOneWayDeliveryTime() != expectedDeliveryTime {
		t.Errorf("Expected GetOneWayDeliveryTime to be %f, but got %f", expectedDeliveryTime, shipment.GetOneWayDeliveryTime())
	}
}
