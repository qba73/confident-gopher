package cars_test

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/cars"
)

func TestNeedsLicenseForVehicles(t *testing.T) {
	t.Parallel()

	vehicles := []string{
		"car",
		"truck",
	}

	for _, vehicle := range vehicles {
		t.Run(vehicle, func(t *testing.T) {
			t.Parallel()

			got := cars.NeedsLicense(vehicle)
			want := true

			if want != got {
				t.Errorf("NeedsLicense(%s) = %v, want %v", vehicle, got, want)
			}
		})
	}
}

func TestDoesntNeedLicenseForVehicles(t *testing.T) {
	t.Parallel()

	vehicles := []string{
		"bike",
		"stroller",
		"e-scooter",
	}

	for _, vehicle := range vehicles {
		t.Run(vehicle, func(t *testing.T) {
			t.Parallel()

			got := cars.NeedsLicense(vehicle)
			want := false

			if want != got {
				t.Errorf("NeedsLicense(%s) = %v, want %v", vehicle, got, want)
			}
		})
	}
}

func TestChooseVehicle(t *testing.T) {
	t.Parallel()

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
			if !cmp.Equal(test.want, got) {
				t.Error(cmp.Diff(test.want, got))
			}

			// The result should be independent of the order in which the choices are given.
			got = cars.ChooseVehicle(test.choice2, test.choice1)
			if !cmp.Equal(test.want, got) {
				t.Error(cmp.Diff(test.want, got))
			}
		})
	}
}

const floatEqualityThreshold = 1e-5

var opt = cmp.Comparer(func(x, y float64) bool {
	delta := math.Abs(x - y)
	mean := math.Abs(x+y) / 2.0
	return delta/mean < floatEqualityThreshold
})

func TestCalculateResellPrice(t *testing.T) {
	t.Parallel()

	for _, tc := range testResellPrices {
		t.Run(tc.name, func(t *testing.T) {
			got := cars.CalculateResellPrice(tc.originalPrice, tc.age)
			if !cmp.Equal(tc.want, got, opt) {
				t.Error(cmp.Diff(tc.want, got))
			}
		})
	}
}

var testResellPrices = []struct {
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

func BenchmarkCalculateResellPrice(b *testing.B) {
	for b.Loop() {
		for _, tc := range testResellPrices {
			cars.CalculateResellPrice(tc.originalPrice, tc.age)
		}
	}
}

func BenchmarkCalculateDiff(b *testing.B) {
	for b.Loop() {
		for _, tc := range testResellPrices {
			_ = cmp.Diff(tc.originalPrice, tc.age) != ""
		}
	}
}

func BenchmarkCheckEquality(b *testing.B) {
	for b.Loop() {
		for _, tc := range testResellPrices {
			cmp.Equal(tc.originalPrice, tc.age)
		}
	}
}
