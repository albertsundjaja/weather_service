package weather

import (
	"fmt"
	"time"
)

var weatherDataCache WeatherDataCache
func init() {
	weatherDataCache = WeatherDataCache{WeatherData{}, time.Now()}
}

func GetWeatherData(city string) (WeatherData, error) {
	var weatherData WeatherData
	var err error

	// stale, fetch new data
	if weatherDataCache.Expiry.Before(time.Now()) {
		weatherData, err = getDataWeatherStack(city)
		if err != nil {
			fmt.Println("Error fetching weather data from weatherstack, using alternative", err)
			weatherData, err = getDataOpenWeather(city)
			if err != nil {
				fmt.Println("Error fetching weather data from openweather", err)
				return weatherData, err
			}
		}
		// update expiry
		weatherDataCache.WeatherData = weatherData
		weatherDataCache.Expiry = time.Now().Add(3 * time.Second)
	} else {
		// serve from cache if not stale
		weatherData = weatherDataCache.WeatherData
	}

	return weatherData, nil
}