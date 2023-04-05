package main

type Package struct {
	Name         string
	Weight       float64
	Distance     float64
	OfferCode    string
	DeliveryTime float64
}

type Shipment struct {
	DeliveryTime float64
	Packages     map[string]*Package
}

func (s *Shipment) addPackages(pkgs []*Package) {
	// assign packages to internal data struct
	for _, p := range pkgs {
		s.Packages[p.Name] = p
	}

	// calculate the delivery time of this shipment, which is
	// max(Packages.DeliveryTime)
	for _, p := range s.Packages {
		if p.DeliveryTime > s.DeliveryTime {
			s.DeliveryTime = p.DeliveryTime
		}
	}
}
