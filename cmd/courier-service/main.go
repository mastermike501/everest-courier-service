package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mastermike501/everest-courier-service/ev_package"
	"github.com/mastermike501/everest-courier-service/fleet"
)

func main() {
	var baseDeliveryCost float64
	var numOfPkgs int
	var err error
	var fleet *fleet.Fleet
	reader := bufio.NewReader(os.Stdin)

	// read delivery cost and num of packages
	for {
		baseDeliveryCost, numOfPkgs, err = readDeliveryCostAndNumOfPkgs(reader)
		if err == nil {
			break
		}
	}

	// read package information
	packages := []*ev_package.Package{}
	for i := 0; i < numOfPkgs; i++ {
		fmt.Printf("Enter package %d information: ", i+1)

		newPackage, err := readPackage(reader, baseDeliveryCost)
		if err != nil {
			i--
			continue
		}

		packages = append(packages, newPackage)
	}

	// read fleet information
	for {
		fleet, err = readFleetInfo(reader)
		if err == nil {
			break
		}
	}

	// delivery time calculations
	for _, p := range packages {
		p.CalculateTimeToDest(fleet.GetMaxSpeed())
	}

	shipments := runKnapsackSolver(packages, fleet)

	FleetSimulation(fleet, shipments)

	fmt.Println()
	fmt.Println("----- Output -----")

	for _, p := range packages {
		p.Println()
	}
}

func runKnapsackSolver(packages []*ev_package.Package, fleet *fleet.Fleet) (shipments []*ev_package.Shipment) {
	remainingPkgs := make([]*ev_package.Package, len(packages))
	copy(remainingPkgs, packages)

	// the "value" would be the weight itself since we are
	// trying to optimize each delivery to carry the heaviest load
	items := []KItem{}
	for _, pkg := range packages {
		item := &KItem{
			weight: pkg.Weight,
			value:  pkg.Weight,
			pkg:    pkg,
		}
		items = append(items, *item)
	}

	// for each shipment do 2 things
	// 1. Calculate the delivery time for each package in shipment and assign it to package
	// 2. Once delivery times for all packages in shipment are calculated, find the maximum.
	// 		The (maximum * 2) is the total time a courier would spend deivering the shipment.

	// iterate over remaining packages and remove packages that have been selected
	// repeat until no more remaining packages
	for len(remainingPkgs) > 0 {
		// solve 0/1 knapsack problem to get selected packages for a shipment
		selected := KnapsackSolver(items, fleet.GetMaxCarriableWeight())

		// for each index returned:
		// 1. get the package
		// 2. calculate the package's delivery time
		// 3. add the package to the list of packages (ie. shipment)
		selectedPkgs := []*ev_package.Package{}
		for _, s := range selected {
			pkg := &remainingPkgs[s]
			selectedPkgs = append(selectedPkgs, *pkg)
		}

		// create a Shipment which takes in the list of packages and generates
		// the total delivery time for the Shipment
		shipment := ev_package.Shipment{
			OneWayDeliveryTime: 0.0,
			Packages:           make(map[string]*ev_package.Package),
		}
		shipment.AddPackages(selectedPkgs)
		shipments = append(shipments, &shipment)

		// remove selected packages
		items = RemoveAtIndexKItem(items, selected)
		remainingPkgs = RemoveAtIndexPackage(remainingPkgs, selected)
	}

	return
}
