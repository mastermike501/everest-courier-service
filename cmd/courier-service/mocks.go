package main

import "bufio"

func mock_readPackages() []Package {
	return []Package{
		{
			Name:      "PKG1",
			Weight:    50.0,
			Distance:  30.0,
			OfferCode: "OFR001",
		},
		{
			Name:      "PKG2",
			Weight:    75.0,
			Distance:  125.0,
			OfferCode: "OFR008",
		},
		{
			Name:      "PKG3",
			Weight:    175.0,
			Distance:  100.0,
			OfferCode: "OFR003",
		},
		{
			Name:      "PKG4",
			Weight:    110.0,
			Distance:  60.0,
			OfferCode: "OFR002",
		},
		{
			Name:      "PKG5",
			Weight:    155.0,
			Distance:  95.0,
			OfferCode: "NA",
		},
	}
}

func mock_readDeliveryCostAndNumOfPkgs(reader *bufio.Reader) (float64, int, error) {
	return 100.0, 5, nil
}

func mock_readFleetInfo(reader *bufio.Reader) (*Fleet, error) {
	return &Fleet{
		NumVehicles:        2,
		MaxSpeed:           70,
		MaxCarriableWeight: 200,
	}, nil
}
