// Package weather provides a tool to print the weather conditions in a certain city.
package weather

var (
	// CurrentCondition represents the current weather.
	CurrentCondition string
	// CurrentLocation represents the current location.
	CurrentLocation  string
)

// Forecast returns a string which shows the current location and its current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
