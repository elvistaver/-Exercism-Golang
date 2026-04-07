//Package weather is the tool that checks the current weather condition of the current location in a city.
package weather

//CurrentCondition stores the value of the current condition of the location.
var CurrentCondition string
 //CurrentLocation stores the value of the current loaction in the city.
var CurrentLocation  string

// Forecast takes in the data for calculation of weather condition to get the current weather conditio.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
