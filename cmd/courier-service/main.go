package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var baseDeliveryCost float64
	var numOfPkgs int
	var err error
	var fleet *Fleet
	reader := bufio.NewReader(os.Stdin)

	// read delivery cost and num of packages
	for {
		baseDeliveryCost, numOfPkgs, err = readDeliveryCostAndNumOfPkgs(reader)
		if err == nil {
			break
		}
	}

	// read package information
	packages := []*Package{}
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
		p.CalculateTimeToDest(fleet.MaxSpeed)
	}

	remainingPkgs := make([]*Package, len(packages))
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
	shipments := []*Shipment{}
	for len(remainingPkgs) > 0 {
		// solve 0/1 knapsack problem to get selected packages for a shipment
		selected := KnapsackSolver(items, fleet.MaxCarriableWeight)

		// for each index returned:
		// 1. get the package
		// 2. calculate the package's delivery time
		// 3. add the package to the list of packages (ie. shipment)
		selectedPkgs := []*Package{}
		for _, s := range selected {
			pkg := &remainingPkgs[s]
			selectedPkgs = append(selectedPkgs, *pkg)
		}

		// create a Shipment which takes in the list of packages and generates
		// the total delivery time for the Shipment
		shipment := Shipment{
			OneWayDeliveryTime: 0.0,
			Packages:           make(map[string]*Package),
		}
		shipment.addPackages(selectedPkgs)
		shipments = append(shipments, &shipment)

		// remove selected packages
		items = RemoveAtIndexKItem(items, selected)
		remainingPkgs = RemoveAtIndexPackage(remainingPkgs, selected)
	}

	FleetSimulation(fleet, shipments)

	fmt.Println()
	fmt.Println("----- Output -----")

	for _, p := range packages {
		p.Println()
	}
}

func readPackage(reader *bufio.Reader, baseDeliveryCost float64) (*Package, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return nil, err
	}

	packageInfo := strings.Split(strings.TrimSpace(input), " ")
	if len(packageInfo) != 4 {
		fmt.Println("The package has missing information. Please try again")
		return nil, fmt.Errorf("missing package information")
	}

	weight, err := strconv.ParseFloat(packageInfo[1], 32)
	if err != nil {
		fmt.Println("An error occured while reading the weight. Please try again", err)
		return nil, err
	}

	distance, err := strconv.ParseFloat(packageInfo[2], 32)
	if err != nil {
		fmt.Println("An error occured while reading the distance. Please try again", err)
		return nil, err
	}

	voucher := getVoucherInfo(packageInfo[3])
	discount, total := GetDiscountAndDeliveryCost(baseDeliveryCost, weight, distance, voucher)

	return &Package{
		Name:      packageInfo[0],
		Weight:    weight,
		Distance:  distance,
		Discount:  discount,
		TotalCost: total,
	}, nil
}

func GetDiscountAndDeliveryCost(baseDeliveryCost, weight, distance float64, v *Voucher) (discount, total float64) {
	grossTotal := baseDeliveryCost + (weight * 10) + (distance * 5)

	// invalid voucher? Return the gross
	if !v.Valid {
		return 0, grossTotal
	}

	// voucher not valid for package parameters? Return the gross
	isVoucherValid := v.ValidForDelivery(distance, weight)
	if !isVoucherValid {
		return 0, grossTotal
	}

	discount = grossTotal * v.Value
	return discount, grossTotal - discount
}

func readDeliveryCostAndNumOfPkgs(reader *bufio.Reader) (float64, int, error) {
	fmt.Print("Enter [Base delivery cost] and [Number of packages]: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return -1, -1, err
	}

	result := strings.Split(strings.TrimSpace(input), " ")

	baseDeliveryCost, err := strconv.ParseFloat(result[0], 32)
	if err != nil {
		fmt.Println("Base delivery cost is an invalid number. Please try again", err)
		return -1, -1, err
	}

	numPkgs, err := strconv.Atoi(result[1])
	if err != nil {
		fmt.Println("Number of packages is an invalid number. Please try again", err)
		return -1, -1, err
	}

	return float64(baseDeliveryCost), numPkgs, nil
}

func readFleetInfo(reader *bufio.Reader) (*Fleet, error) {
	fmt.Print("Enter fleet info: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return nil, err
	}

	result := strings.Split(strings.TrimSpace(input), " ")

	if len(result) != 3 {
		fmt.Println("Not enough information provided. Please try again")
		return nil, fmt.Errorf("not enough information provided")
	}

	numVehicles, err := strconv.Atoi(result[0])
	if err != nil {
		fmt.Println("Number of vehicles is an invalid number. Please try again", err)
		return nil, err
	}

	maxSpeed, err := strconv.Atoi(result[1])
	if err != nil {
		fmt.Println("Max speed is an invalid number. Please try again", err)
		return nil, err
	}

	maxCarriableWeight, err := strconv.ParseFloat(result[2], 64)
	if err != nil {
		fmt.Println("Max carriable weight is an invalid number. Please try again", err)
		return nil, err
	}

	return &Fleet{
		NumVehicles:        numVehicles,
		MaxSpeed:           maxSpeed,
		MaxCarriableWeight: maxCarriableWeight,
	}, nil
}
