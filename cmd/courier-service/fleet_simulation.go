package main

import (
	"fmt"

	"github.com/mastermike501/everest-courier-service/ev_package"
	"github.com/mastermike501/everest-courier-service/fleet"
)

// Calculate package delivery times by performing discrete Simulate
func FleetSimulation(f *fleet.Fleet, shipments []*ev_package.Shipment) {
	currentTime := 0.0

	// create vehicles
	// fleet.
	vehicles := []*fleet.Vehicle{}
	for i := 0; i < f.GetNumVehicles(); i++ {
		v := fleet.NewVehicle(fmt.Sprint(i+1), 0.0)
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
		for _, p := range s.GetPackages() {
			p.DeliveryTime = currentTime + p.TimeToDest
		}

		// set the return time for selected vehicle
		selectedVehicle.SetReturnTime((s.GetOneWayDeliveryTime() * 2) + currentTime)

		shipmentIdx++

		// once all shipments are settled, break out of the loop
		// all update code after this if statement is no longer relevant
		if shipmentIdx == len(shipments) {
			break
		}
	}
}

func getMinimumReturnTime(vehicles []*fleet.Vehicle) float64 {
	minRetTime := vehicles[0].GetReturnTime()

	for _, v := range vehicles {
		if v.GetReturnTime() < minRetTime {
			minRetTime = v.GetReturnTime()
		}
	}

	return minRetTime
}

func updateVehicleReturnTimes(vehicles []*fleet.Vehicle, curTime float64) {
	for _, v := range vehicles {
		v.UpdateReturnTime(curTime)
	}
}

func selectFreeVehicle(vehicles []*fleet.Vehicle) *fleet.Vehicle {
	for _, v := range vehicles {
		if v.GetReturnTime() == 0.0 {
			return v
		}
	}

	return nil
}
