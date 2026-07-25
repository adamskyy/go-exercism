package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (tu TemperatureUnit) String() string {
    if tu == Celsius { 
        return fmt.Sprintf("°C")
    } else if tu == Fahrenheit {
        return fmt.Sprintf("°F")
    }
    return ""
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (tem Temperature) String() string {
    return fmt.Sprintf("%d %s", tem.degree, tem.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (speedUnit SpeedUnit) String() string {
    if speedUnit == KmPerHour { 
        return fmt.Sprintf("km/h")
    } else if speedUnit == MilesPerHour {
        return fmt.Sprintf("mph")
    }
    return ""
}


type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (speed Speed) String() string {
    return fmt.Sprintf("%d %s", speed.magnitude, speed.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (metData MeteorologyData) String() string {
    return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", metData.location, metData.temperature, metData.windDirection, metData.windSpeed, metData.humidity)
}
