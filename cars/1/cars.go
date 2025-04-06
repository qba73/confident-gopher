package cars

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is needed
// to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

func ChooseCarByLexicographicalOrder(car1, car2 string) string {
	if strings.Compare(car1, car2) < 0 {
		return car1
	}
	return car2
}

// ChooseVehicle recommends a vehicle for selection.
// It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	chosenCar := ChooseCarByLexicographicalOrder(option1, option2)
	return fmt.Sprintf("%s is clearly the better choice.", chosenCar)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 80 / 100
	} else if age >= 10 {
		return originalPrice * 50 / 100
	} else {
		return originalPrice * 70 / 100
	}
}
