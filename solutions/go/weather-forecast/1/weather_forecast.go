// Package weather represents blah blah blah.
package weather

var (
    // CurrentCondition represents current weather condition.
	CurrentCondition string
    // CurrentLocation represents current location city.
	CurrentLocation  string
)

// Forecast function returns strings representing .
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
