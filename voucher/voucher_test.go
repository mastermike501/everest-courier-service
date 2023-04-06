package voucher

import (
	"testing"
)

func TestVoucher_GetValue(t *testing.T) {
	tests := []struct {
		voucher Voucher
		value   float64
	}{
		{Voucher{value: 0.1}, 0.1},
		{Voucher{value: 0.07}, 0.07},
		{Voucher{value: 0.05}, 0.05},
	}

	for _, tt := range tests {
		if tt.voucher.GetValue() != tt.value {
			t.Errorf("Expected voucher value %f but got %f", tt.value, tt.voucher.GetValue())
		}
	}
}

func TestVoucher_IsValid(t *testing.T) {
	tests := []struct {
		voucher Voucher
		valid   bool
	}{
		{Voucher{valid: true}, true},
		{Voucher{valid: false}, false},
	}

	for _, tt := range tests {
		if tt.voucher.IsValid() != tt.valid {
			t.Errorf("Expected voucher validity %t but got %t", tt.valid, tt.voucher.IsValid())
		}
	}
}

func TestGetVoucherInfo(t *testing.T) {
	tests := []struct {
		code        string
		description string
		value       float64
		valid       bool
	}{
		{"OFR001", "10% Discount", 0.1, true},
		{"OFR002", "7% Discount", 0.07, true},
		{"OFR003", "5% Discount", 0.05, true},
		{"OFR004", "No Discount", 0, false},
	}

	for _, tt := range tests {
		voucher := GetVoucherInfo(tt.code)

		if voucher.description != tt.description {
			t.Errorf("Expected voucher description %s but got %s", tt.description, voucher.description)
		}

		if voucher.value != tt.value {
			t.Errorf("Expected voucher value %f but got %f", tt.value, voucher.value)
		}

		if voucher.valid != tt.valid {
			t.Errorf("Expected voucher validity %t but got %t", tt.valid, voucher.valid)
		}
	}
}

func TestVoucher_ValidForDelivery(t *testing.T) {
	testCases := []struct {
		voucherCode  string
		distance     float64
		weight       float64
		expectedBool bool
	}{
		{voucherCode: "OFR001", distance: 100, weight: 100, expectedBool: true},
		{voucherCode: "OFR001", distance: 100, weight: 70, expectedBool: true},
		{voucherCode: "OFR001", distance: 199, weight: 200, expectedBool: true},
		{voucherCode: "OFR001", distance: 200, weight: 150, expectedBool: false},
		{voucherCode: "OFR002", distance: 49, weight: 200, expectedBool: false},
		{voucherCode: "OFR002", distance: 50, weight: 99, expectedBool: false},
		{voucherCode: "OFR002", distance: 50, weight: 100, expectedBool: true},
		{voucherCode: "OFR002", distance: 150, weight: 250, expectedBool: true},
		{voucherCode: "OFR002", distance: 151, weight: 250, expectedBool: false},
		{voucherCode: "OFR003", distance: 49, weight: 150, expectedBool: false},
		{voucherCode: "OFR003", distance: 50, weight: 9, expectedBool: false},
		{voucherCode: "OFR003", distance: 50, weight: 10, expectedBool: true},
		{voucherCode: "OFR003", distance: 250, weight: 150, expectedBool: true},
		{voucherCode: "OFR003", distance: 251, weight: 150, expectedBool: false},
		{voucherCode: "invalid", distance: 100, weight: 100, expectedBool: false},
	}

	for _, tc := range testCases {
		voucher := GetVoucherInfo(tc.voucherCode)
		actualBool := voucher.ValidForDelivery(tc.distance, tc.weight)
		if actualBool != tc.expectedBool {
			t.Errorf("Expected ValidForDelivery() to return %v for voucher %s with distance %f and weight %f; got %v", tc.expectedBool, tc.voucherCode, tc.distance, tc.weight, actualBool)
		}
	}
}
