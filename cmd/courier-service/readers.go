package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/mastermike501/everest-courier-service/fleet"
)

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

	voucher := GetVoucherInfo(packageInfo[3])
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

func readFleetInfo(reader *bufio.Reader) (*fleet.Fleet, error) {
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

	return &fleet.Fleet{
		NumVehicles:        numVehicles,
		MaxSpeed:           maxSpeed,
		MaxCarriableWeight: maxCarriableWeight,
	}, nil
}
