package voucher

type Voucher struct {
	description      string
	value            float64
	valid            bool
	ValidForDelivery func(distance, weight float64) bool
}

func (v *Voucher) GetValue() float64 {
	return v.value
}

func (v *Voucher) IsValid() bool {
	return v.valid
}

var vouchers = map[string]Voucher{
	"OFR001": {
		description: "10% Discount",
		value:       0.1,
		valid:       true,
		ValidForDelivery: func(distance, weight float64) bool {
			return distance < 200 && (weight >= 70 && weight <= 200)
		},
	},
	"OFR002": {
		description: "7% Discount",
		value:       0.07,
		valid:       true,
		ValidForDelivery: func(distance, weight float64) bool {
			return (distance >= 50 && distance <= 150) && (weight >= 100 && weight <= 250)
		},
	},
	"OFR003": {
		description: "5% Discount",
		value:       0.05,
		valid:       true,
		ValidForDelivery: func(distance, weight float64) bool {
			return (distance >= 50 && distance <= 250) && (weight >= 10 && weight <= 150)
		},
	},
}

func GetVoucherInfo(code string) *Voucher {
	if voucher, ok := vouchers[code]; ok {
		return &voucher
	}

	return &Voucher{
		description: "No Discount",
		value:       0,
		valid:       false,
		ValidForDelivery: func(distance, weight float64) bool {
			return false
		},
	}
}
