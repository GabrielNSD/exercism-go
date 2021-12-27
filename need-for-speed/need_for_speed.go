package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
		battery:      100,
		speed:        speed,
		batteryDrain: batteryDrain,
		distance:     0,
	}
	return car
}

// TODO: define the 'Track' type struct

type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	track := Track{
		distance: distance,
	}
	return track
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	newBattery := car.battery - car.batteryDrain
	newDistance := car.distance + car.speed
	if newBattery < 0 {
		return car
	} else {
		return Car{speed: car.speed, batteryDrain: car.batteryDrain, battery: newBattery, distance: newDistance}
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	maxDistance := (car.battery / car.batteryDrain) * car.speed
	if maxDistance >= track.distance {
		return true
	} else {
		return false
	}
}
