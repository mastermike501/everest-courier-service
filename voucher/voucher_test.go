package voucher

import (
	"reflect"
	"testing"
)

func TestGetVoucherInfo(t *testing.T) {
	t.Run("Valid voucher code OFR001", func(t *testing.T) {
		tests := []struct {
			name string
			code string
			want *Voucher
		}{
			{
				name: "Test case 1",
				code: "OFR001",
				want: &Voucher{
					description: "10% Discount",
					value:       0.1,
					ValidForDelivery: func(distance, weight float64) bool {
						return distance < 200 && (weight >= 70 && weight <= 200)
					},
				},
			},
			{
				name: "Test case 2",
				code: "OFR003",
				want: &Voucher{
					description: "5% Discount",
					value:       0.05,
					ValidForDelivery: func(distance, weight float64) bool {
						return (distance >= 50 && distance <= 250) && (weight >= 10 && weight <= 150)
					},
				},
			},
			{
				name: "Test case 3 - invalid code",
				code: "OFR004",
				want: &Voucher{
					description: "No Discount",
					value:       0,
					ValidForDelivery: func(distance, weight float64) bool {
						return false
					},
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := GetVoucherInfo(tt.code); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("getVoucherInfo() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("Invalid voucher code", func(t *testing.T) {
		voucher := GetVoucherInfo("INVALID")
		if voucher.description != "No Discount" {
			t.Errorf("Expected voucher.Description to be 'No Discount', but got %v", voucher.description)
		}
		if voucher.value != 0 {
			t.Errorf("Expected voucher.Value to be 0, but got %v", voucher.value)
		}
		if !voucher.ValidForDelivery(150, 150) {
			t.Errorf("Expected voucher.ValidForDelivery to return true for any distance and weight")
		}
	})
}