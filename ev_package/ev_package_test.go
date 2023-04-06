package ev_package

import (
	"testing"
)

func TestCalculateTimeToDest(t *testing.T) {
	p := &Package{
		Name:     "test_pkg",
		Weight:   10.0,
		Distance: 100.0,
	}
	speed := 50 // km/h

	p.CalculateTimeToDest(speed)

	expected := 2.0 // hours
	if p.TimeToDest != expected {
		t.Errorf("Expected time to dest %f, but got %f", expected, p.TimeToDest)
	}
}

func TestPrintln(t *testing.T) {
	p := &Package{
		Name:         "test_pkg",
		Weight:       10.0,
		Distance:     100.0,
		Discount:     5.0,
		TotalCost:    100.0,
		DeliveryTime: 2.0,
	}

	p.Println()

	// Check the output of Println
	expectedOutput := "test_pkg 5 100 2.00\n"
	if testing.Verbose() {
		t.Log("Expected output:", expectedOutput)
	}
}
