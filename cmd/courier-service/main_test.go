package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadPackage(t *testing.T) {
	testCases := []struct {
		input    string
		expected Package
		wantErr  bool
	}{
		{
			input:    "PackageA 20 100 OFR001\n",
			expected: Package{Name: "PackageA", Weight: 20, Distance: 100, OfferCode: "OFR001"},
			wantErr:  false,
		},
		{
			input:    "PackageB 30 150 OFR002\n",
			expected: Package{Name: "PackageB", Weight: 30, Distance: 150, OfferCode: "OFR002"},
			wantErr:  false,
		},
		{
			input:    "PackageC 40 200 OFR003\n",
			expected: Package{Name: "PackageC", Weight: 40, Distance: 200, OfferCode: "OFR003"},
			wantErr:  false,
		},
		{
			input:    "PackageD 50 250\n",
			expected: Package{},
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		reader := strings.NewReader(tc.input)
		bufReader := bufio.NewReader(reader)
		pkg, err := readPackage(bufReader)
		if tc.wantErr && err == nil {
			t.Errorf("Expected error for input: %s", tc.input)
		}
		if !tc.wantErr && err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !tc.wantErr && *pkg != tc.expected {
			t.Errorf("Expected: %v but got: %v", tc.expected, *pkg)
		}
	}
}

func TestCalculateDeliveryCost(t *testing.T) {
	type args struct {
		baseDeliveryCost float64
		pkg              *Package
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test case 1",
			args: args{
				baseDeliveryCost: 100,
				pkg: &Package{
					Name:      "Test package",
					Weight:    10,
					Distance:  50,
					OfferCode: "",
				},
			},
			want: 650,
		},
		{
			name: "Test case 2",
			args: args{
				baseDeliveryCost: 50,
				pkg: &Package{
					Name:      "Test package",
					Weight:    5,
					Distance:  200,
					OfferCode: "",
				},
			},
			want: 300,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateDeliveryCost(tt.args.baseDeliveryCost, tt.args.pkg); got != tt.want {
				t.Errorf("calculateDeliveryCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadDeliveryCostAndNumOfPkgs(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		input := "25.50 4\n"
		expectedBaseCost := 25.50
		expectedNumOfPkgs := 4

		reader := bufio.NewReader(strings.NewReader(input))

		cost, num, err := readDeliveryCostAndNumOfPkgs(reader)
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}
		if cost != expectedBaseCost {
			t.Errorf("Expected base delivery cost to be %.2f, but got %.2f", expectedBaseCost, cost)
		}
		if num != expectedNumOfPkgs {
			t.Errorf("Expected number of packages to be %d, but got %d", expectedNumOfPkgs, num)
		}
	})

	t.Run("invalid base cost", func(t *testing.T) {
		input := "invalid 4\n"

		reader := bufio.NewReader(strings.NewReader(input))

		_, _, err := readDeliveryCostAndNumOfPkgs(reader)
		if err == nil {
			t.Errorf("Expected error, but got none")
		}
	})

	t.Run("invalid number of packages", func(t *testing.T) {
		input := "25.50 invalid\n"

		reader := bufio.NewReader(strings.NewReader(input))

		_, _, err := readDeliveryCostAndNumOfPkgs(reader)
		if err == nil {
			t.Errorf("Expected error, but got none")
		}
	})
}
