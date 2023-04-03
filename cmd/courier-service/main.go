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
	reader := bufio.NewReader(os.Stdin)

	// read delivery cost and num of packages
	for {
		baseDeliveryCost, numOfPkgs, err = readDeliveryCostAndNumOfPkgs(reader)
		if err == nil {
			break
		}
	}

	packages := []Package{}

	// read package information
	for i := 0; i < numOfPkgs; i++ {
		fmt.Printf("Enter package %d information: ", i+1)

		newPackage, err := readPackage(reader)
		if err != nil {
			i--
			continue
		}

		packages = append(packages, *newPackage)
	}

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

func calculateDeliveryCost(baseDeliveryCost float64, pkg *Package) float64 {
	return baseDeliveryCost + (pkg.Weight * 10) + (pkg.Distance * 5)
}
