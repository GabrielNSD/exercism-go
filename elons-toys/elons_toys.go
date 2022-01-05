package elon

import "fmt"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	newDistance := car.distance + car.speed
	newBattery := car.battery - car.batteryDrain

	if newBattery >= 0 {
		car.distance = newDistance
		car.battery = newBattery
	}
}

// TODO: define the 'DisplayDistance() string' method
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car Car) CanFinish(trackDistance int) bool {
	maxDistance := (car.battery / car.batteryDrain) * car.speed
	if maxDistance >= trackDistance {
		return true
	} else {
		return false
	}

}
