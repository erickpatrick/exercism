// Package weather provides tools to get the
// weather forecast for a given city.
package weather

// CurrentCondition holds the current weather condition for a given city.
var CurrentCondition string

// CurrentLocation holds the current city which we're weather forecasting.
var CurrentLocation string

// Forecast returns a string with current city and weather condition for that given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
