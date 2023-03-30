package main

type Voucher struct {
	Description      string
	ValidForDelivery func(distance, weight int) bool
}

var vouchers = map[string]Voucher{
	"OFR001": {
		Description: "10% Discount",
		ValidForDelivery: func(distance, weight int) bool {
			if distance < 200 && (weight >= 70 && weight <= 200) {
				return true
			}
			return false
		},
	},
	"OFR002": {
		Description: "7% Discount",
		ValidForDelivery: func(distance, weight int) bool {
			if (distance >= 50 && distance <= 150) && (weight >= 100 && weight <= 250) {
				return true
			}
			return false
		},
	},
	"OFR003": {
		Description: "5% Discount",
		ValidForDelivery: func(distance, weight int) bool {
			if (distance >= 50 && distance <= 250) && (weight >= 10 && weight <= 150) {
				return true
			}
			return false
		},
	},
}
