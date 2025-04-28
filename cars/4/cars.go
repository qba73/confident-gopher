package cars

// NeedsLicense determines whether a license is needed
// to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

func ChooseCar(car1, car2 string) string {
	if car1 < car2 {
		return car1
	}
	return car2
}

// ChooseVehicle recommends a vehicle for selection.
// It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	car := ChooseCar(option1, option2)
	return car + " is clearly the better choice."
}

// CalculateResellPrice calculates how much
// a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 80 / 100
	}
	if age < 10 {
		return originalPrice * 70 / 100
	}
	return originalPrice * 50 / 100
}
