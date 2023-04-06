package main

type Voucher struct {
	Description      string
	Value            float64
	Valid            bool
	ValidForDelivery func(distance, weight float64) bool
}

var vouchers = map[string]Voucher{
	"OFR001": {
		Description: "10% Discount",
		Value:       0.1,
		Valid:       true,
		ValidForDelivery: func(distance, weight float64) bool {
			return distance < 200 && (weight >= 70 && weight <= 200)
		},
	},
	"OFR002": {
		Description: "7% Discount",
		Value:       0.07,
		Valid:       true,
		ValidForDelivery: func(distance, weight float64) bool {
			return (distance >= 50 && distance <= 150) && (weight >= 100 && weight <= 250)
		},
	},
	"OFR003": {
		Description: "5% Discount",
		Value:       0.05,
		Valid:       true,
		ValidForDelivery: func(distance, weight float64) bool {
			return (distance >= 50 && distance <= 250) && (weight >= 10 && weight <= 150)
		},
	},
}

func getVoucherInfo(code string) *Voucher {
	if voucher, ok := vouchers[code]; ok {
		return &voucher
	}

	return &Voucher{
		Description: "No Discount",
		Value:       0,
		Valid:       false,
		ValidForDelivery: func(distance, weight float64) bool {
			return false
		},
	}
}
