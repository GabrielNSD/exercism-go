// Package weather contains a function called Forecast.
package weather

// CurrentCondition variable stores Current Condition as a string.
var CurrentCondition string

// CurrentLocation stores Current Location as a string.
var CurrentLocation string

// Forecast returns wheater condition in given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
