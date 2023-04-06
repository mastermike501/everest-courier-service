package main

import (
	"fmt"
	"math"
)

type Package struct {
	Name      string
	Weight    float64
	Distance  float64
	Discount  float64
	TotalCost float64

	TimeToDest   float64 // exact time taken from depot to dest based on vehicle speed
	DeliveryTime float64 // TimeToDest + current time
}

// pkg_id1 discount1 total_cost1 estimated_delivery_time1_in_hours
func (p *Package) Println() {
	fmt.Printf("%s %.0f %.0f %.2f\n", p.Name, p.Discount, p.TotalCost, p.DeliveryTime)
}

func (p *Package) CalculateTimeToDest(speed int) {
	t := p.Distance / float64(speed)
	p.TimeToDest = math.Floor(t*100) / 100 // rounds down to 2 dec places
}

type Shipment struct {
	OneWayDeliveryTime float64
	Packages           map[string]*Package
}

func (s *Shipment) AddPackages(pkgs []*Package) {
	// assign packages to internal data struct
	for _, p := range pkgs {
		s.Packages[p.Name] = p
	}

	// calculate the delivery time of this shipment, which is
	// max(Packages.DeliveryTime)
	for _, p := range s.Packages {
		if p.TimeToDest > s.OneWayDeliveryTime {
			s.OneWayDeliveryTime = p.TimeToDest
		}
	}
}
