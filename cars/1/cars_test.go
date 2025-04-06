package cars_test

import (
	"math"
	"testing"

	"github.com/qba73/cars"
)

const floatEqualityThreshold = 1e-5

func TestNeedsLicense(t *testing.T) {
	tests := []struct {
		name string
		kind string
		want bool
	}{
		{
			name: "need a license for a car",
			kind: "car",
			want: true,
		},
		{
			name: "need a license for a truck",
			kind: "truck",
			want: true,
		},
		{
			name: "does not need a license for a bike",
			kind: "bike",
			want: false,
		},
		{
			name: "does not need a license for a stroller",
			kind: "stroller",
			want: false,
		},
		{
			name: "does not need a license for a e-scooter",
			kind: "e-scooter",
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := cars.NeedsLicense(test.kind)
			if got != test.want {
				t.Errorf(
					"NeedsLicense(%q) = %t, want %t",
					test.kind,
					got,
					test.want)
			}
		})
	}
}

func TestChooseVehicle(t *testing.T) {
	tests := []struct {
		name    string
		choice1 string
		choice2 string
		want    string
	}{
		{
			name:    "chooses Bugatti over Ford",
			choice1: "Bugatti Veyron",
			choice2: "Ford Pinto",
			want:    "Bugatti Veyron is clearly the better choice.",
		}, {
			name:    "chooses Chery over Kia",
			choice1: "Chery EQ",
			choice2: "Kia Niro Elektro",
			want:    "Chery EQ is clearly the better choice.",
		}, {
			name:    "chooses Ford Focus over Ford Pinto",
			choice1: "Ford Focus",
			choice2: "Ford Pinto",
			want:    "Ford Focus is clearly the better choice.",
		}, {
			name:    "chooses 2018 over 2020",
			choice1: "2018 Bergamont City",
			choice2: "2020 Gazelle Medeo",
			want:    "2018 Bergamont City is clearly the better choice.",
		}, {
			name:    "chooses Bugatti over ford",
			choice1: "Bugatti Veyron",
			choice2: "ford",
			want:    "Bugatti Veyron is clearly the better choice.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := cars.ChooseVehicle(test.choice1, test.choice2)
			if got != test.want {
				t.Errorf(
					"ChooseVehicle(%q, %q) = %q, want %q",
					test.choice1,
					test.choice2,
					got,
					test.want)
			}

			// The result should be independent of the order in which the choices are given.
			got = cars.ChooseVehicle(test.choice2, test.choice1)
			if got != test.want {
				t.Errorf(
					"ChooseVehicle(%q, %q) = %q, want %q",
					test.choice2,
					test.choice1,
					got,
					test.want)
			}
		})
	}
}

func TestCalculateResellPrice(t *testing.T) {
	tests := []struct {
		name          string
		originalPrice float64
		age           float64
		want          float64
	}{
		{
			name:          "price is reduced to 80% for age below 3",
			originalPrice: 40000,
			age:           2,
			want:          32000,
		},
		{
			name:          "price is reduced to 80% for age below 3",
			originalPrice: 40000,
			age:           2.5,
			want:          32000,
		},
		{
			name:          "price is reduced to 70% for age 3",
			originalPrice: 40000,
			age:           3,
			want:          28000,
		},
		{
			name:          "price is reduced to 70% for age 7",
			originalPrice: 40000,
			age:           7,
			want:          28000,
		},
		{
			name:          "price is reduced to 50% for age 10",
			originalPrice: 25000,
			age:           10,
			want:          12500,
		},
		{
			name:          "price is reduced to 50% for age 11",
			originalPrice: 50000,
			age:           11,
			want:          25000,
		},
		{
			name:          "float price is reduced to 70% for age 8,",
			originalPrice: 39000.000001,
			age:           8,
			want:          27300.0000007,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := cars.CalculateResellPrice(test.originalPrice, test.age)
			if !floatingPointEquals(got, test.want) {
				t.Errorf(
					"CalculateResellPrice(%v, %v) = %v, want %v",
					test.originalPrice,
					test.age,
					got,
					test.want)
			}
		})
	}
}

func floatingPointEquals(got, want float64) bool {
	absoluteDifferenceBelowThreshold := math.Abs(got-want) <= floatEqualityThreshold
	relativeDifferenceBelowThreshold := math.Abs(got-want)/(math.Abs(got)+math.Abs(want)) <= floatEqualityThreshold
	return absoluteDifferenceBelowThreshold || relativeDifferenceBelowThreshold
}
