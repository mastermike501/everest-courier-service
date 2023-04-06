package main

import "fmt"

// Calculate package delivery times by performing discrete Simulate
func FleetSimulation(fleet *Fleet, shipments []*Shipment) {
	currentTime := 0.0

	// create vehicles
	vehicles := []*Vehicle{}
	for i := 0; i < fleet.NumVehicles; i++ {
		v := &Vehicle{
			Name:       fmt.Sprint(i + 1),
			ReturnTime: 0.0,
		}
		vehicles = append(vehicles, v)
	}

	// the current time is incremented when there are all vehicles are busy
	// the next current time is the minimum return time of all vehicles
	shipmentIdx := 0
	for {
		s := shipments[shipmentIdx]

		// 1. If no available vehicles, increment the current time and update the
		// 	vehicle return times
		// 2. If a vehicle is available, do not increment current time and update
		//	selected vehicle with information, then inc shipmentIdx
		selectedVehicle := selectFreeVehicle(vehicles)
		if selectedVehicle == nil {
			// This line updates current time only if all vehicles are busy.
			// If there is any vehicle that is available, it will have a ReturnTime
			// of 0. So getMinimumReturnTime will return 0 or the next value to inc to
			currentTime = currentTime + getMinimumReturnTime(vehicles)

			// update vehicle return times
			updateVehicleReturnTimes(vehicles, currentTime)
			continue
		}

		// update delivery times for packages
		for _, p := range s.Packages {
			p.DeliveryTime = currentTime + p.TimeToDest
		}

		// set the return time for selected vehicle
		selectedVehicle.ReturnTime = (s.OneWayDeliveryTime * 2) + currentTime

		shipmentIdx++

		// once all shipments are settled, break out of the loop
		// all update code after this if statement is no longer relevant
		if shipmentIdx == len(shipments) {
			break
		}
	}
}

func getMinimumReturnTime(vehicles []*Vehicle) float64 {
	minRetTime := vehicles[0].ReturnTime

	for _, v := range vehicles {
		if v.ReturnTime < minRetTime {
			minRetTime = v.ReturnTime
		}
	}

	return minRetTime
}

func updateVehicleReturnTimes(vehicles []*Vehicle, curTime float64) {
	for _, v := range vehicles {
		v.updateReturnTime(curTime)
	}
}

func selectFreeVehicle(vehicles []*Vehicle) *Vehicle {
	for _, v := range vehicles {
		if v.ReturnTime == 0.0 {
			return v
		}
	}

	return nil
}
