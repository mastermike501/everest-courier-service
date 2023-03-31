package main

type Voucher struct {
	Description      string
	ValidForDelivery func(distance, weight float32) bool
}

var vouchers = map[string]Voucher{
	"OFR001": {
		Description: "10% Discount",
		ValidForDelivery: func(distance, weight float32) bool {
			return distance < 200 && (weight >= 70 && weight <= 200)
		},
	},
	"OFR002": {
		Description: "7% Discount",
		ValidForDelivery: func(distance, weight float32) bool {
			return (distance >= 50 && distance <= 150) && (weight >= 100 && weight <= 250)
		},
	},
	"OFR003": {
		Description: "5% Discount",
		ValidForDelivery: func(distance, weight float32) bool {
			return (distance >= 50 && distance <= 250) && (weight >= 10 && weight <= 150)
		},
	},
}

func getVoucherInfo(code string) *Voucher {
	if voucher, ok := vouchers[code]; ok {
		return &voucher
	}

	return nil
}
