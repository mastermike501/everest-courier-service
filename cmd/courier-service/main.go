package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Fleet struct {
	NumVehicles        int
	MaxSpeed           int
	MaxCarriableWeight float64
}

func main() {
	var baseDeliveryCost float64
	var numOfPkgs int
	var err error
	var fleet *Fleet
	reader := bufio.NewReader(os.Stdin)
	baseDeliveryCost, numOfPkgs, err = mock_readDeliveryCostAndNumOfPkgs(reader)
	packages := mock_readPackages()
	fleet, err = mock_readFleetInfo(reader)

	if err != nil {
		os.Exit(1)
	}

	fmt.Print(numOfPkgs)

	// read delivery cost and num of packages
	// for {
	// 	baseDeliveryCost, numOfPkgs, err = readDeliveryCostAndNumOfPkgs(reader)
	// 	if err == nil {
	// 		break
	// 	}
	// }

	// packages := []Package{}

	// // read package information
	// for i := 0; i < numOfPkgs; i++ {
	// 	fmt.Printf("Enter package %d information: ", i+1)

	// 	newPackage, err := readPackage(reader)
	// 	if err != nil {
	// 		i--
	// 		continue
	// 	}

	// 	packages = append(packages, *newPackage)
	// }

	// read fleet information
	// for {
	// 	fleet, err = readFleetInfo(reader)
	// 	if err == nil {
	// 		break
	// 	}
	// }

	for _, pkg := range packages {
		deliveryCost := calculateDeliveryCost(baseDeliveryCost, &pkg)
		voucherInfo := getVoucherInfo(pkg.OfferCode)
		discount := 0.0
		totalCost := deliveryCost

		isVoucherValid := voucherInfo.ValidForDelivery(pkg.Distance, pkg.Weight)
		if isVoucherValid {
			discount = deliveryCost * voucherInfo.Value
			totalCost = deliveryCost - discount
		}

		fmt.Printf("%s $%.2f $%.2f", pkg.Name, discount, totalCost)
		fmt.Println()
	}

	remainingPkgs := make([]Package, len(packages))
	copy(remainingPkgs, packages)

	// the "value" would be the weight itself since we are
	// trying to optimize each delivery to carry the heaviest load
	items := []KItem{}
	for _, pkg := range packages {
		item := &KItem{
			weight: pkg.Weight,
			value:  pkg.Weight,
			pkg:    &pkg,
		}
		items = append(items, *item)
	}

	// iterate over items array and remove selected items
	// repeat until items array is empty
	shipment := 1
	for len(remainingPkgs) > 0 {
		selected := KnapsackSolver(items, fleet.MaxCarriableWeight)

		fmt.Printf("Shipment %v\n", shipment)
		for _, s := range selected {
			fmt.Printf("%s ", remainingPkgs[s].Name)
		}
		fmt.Println()
		fmt.Println()

		// remove selected items
		items = RemoveAtIndexKItem(items, selected)
		remainingPkgs = RemoveAtIndexPackage(remainingPkgs, selected)

		shipment++
	}

}

func readPackage(reader *bufio.Reader) (*Package, error) {
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

	return &Package{
		Name:      packageInfo[0],
		Weight:    float64(weight),
		Distance:  float64(distance),
		OfferCode: packageInfo[3],
	}, nil
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

	numVehicles, err := strconv.Atoi(result[0])
	if err != nil {
		fmt.Println("Number of vehicles is an invalid number. Please try again", err)
		return nil, err
	}

	maxSpeed, err := strconv.Atoi(result[0])
	if err != nil {
		fmt.Println("Max speed is an invalid number. Please try again", err)
		return nil, err
	}

	maxCarriableWeight, err := strconv.ParseFloat(result[0], 64)
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

func calculateDeliveryCost(baseDeliveryCost float64, pkg *Package) float64 {
	return baseDeliveryCost + (pkg.Weight * 10) + (pkg.Distance * 5)
}
