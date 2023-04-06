package ev_package

type Shipment struct {
	oneWayDeliveryTime float64
	packages           map[string]*Package
}

func NewShipment() *Shipment {
	return &Shipment{
		oneWayDeliveryTime: 0.0,
		packages:           make(map[string]*Package),
	}
}

func (s *Shipment) GetPackages() map[string]*Package {
	return s.packages
}

func (s *Shipment) GetOneWayDeliveryTime() float64 {
	return s.oneWayDeliveryTime
}

func (s *Shipment) AddPackages(pkgs []*Package) {
	// assign packages to internal data struct
	for _, p := range pkgs {
		s.packages[p.Name] = p
	}

	// calculate the delivery time of this shipment, which is
	// max(Packages.DeliveryTime)
	for _, p := range s.packages {
		if p.TimeToDest > s.oneWayDeliveryTime {
			s.oneWayDeliveryTime = p.TimeToDest
		}
	}
}
