package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	baseDeliveryCost, numOfPkgs, err := readDeliveryCostAndNumOfPkgs()
	if err != nil {
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	packages := []Package{}

	// read package information
	for i := 0; i < numOfPkgs; i++ {
		fmt.Printf("Enter package %d information: ", i+1)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			i--
			continue
		}

		newPackage, err := readPackage(input)
		if err != nil {
			i--
			continue
		}

		packages = append(packages, *newPackage)
	}

	for _, pkg := range packages {
		deliveryCost := calculateDeliveryCost(baseDeliveryCost, &pkg)
		voucherInfo := getVoucherInfo(pkg.OfferCode)
		fmt.Printf("Delivery cost: $%.2f\n", deliveryCost)
		fmt.Printf("Voucher: %s\n", voucherInfo.Description)
	}
}

func readPackage(input string) (*Package, error) {
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
		Weight:    float32(weight),
		Distance:  float32(distance),
		OfferCode: packageInfo[3],
	}, nil
}

func readDeliveryCostAndNumOfPkgs() (float32, int, error) {
	reader := bufio.NewReader(os.Stdin)

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

	return float32(baseDeliveryCost), numPkgs, nil
}

func calculateDeliveryCost(baseDeliveryCost float32, pkg *Package) float32 {
	return baseDeliveryCost + (pkg.Weight * 10) + (pkg.Distance * 6)
}
