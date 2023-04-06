package ev_package

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
