package weather

import (
	"encoding/json"
	"fmt"
	"github.com/albertsundjaja/weather_service/config"
	"net/http"
	"time"
)

type WeatherData struct {
	WindSpeed          float32 `json:"wind_speed"`
	TemperatureDegrees float32 `json:"temperature_degrees"`
}

type WeatherStackData struct {
	Current struct {
		Temperature         float32      `json:"temperature"`
		WindSpeed           float32      `json:"wind_speed"`
	} `json:"current"`
}

type OpenWeatherData struct {
	Main struct {
		Temp      float32 `json:"temp"`
	} `json:"main"`
	Wind       struct {
		Speed float32 `json:"speed"`
	} `json:"wind"`
}

type WeatherDataCache struct {
	WeatherData
	Expiry time.Time
}

func getDataWeatherStack(city string) (WeatherData, error) {
	resp, err := http.Get(fmt.Sprintf("http://api.weatherstack.com/current?access_key=%s&query=%s", config.WEATHER_STACK_API_KEY, city))
	if err != nil {
		return WeatherData{}, err
	}

	var data WeatherStackData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return WeatherData{}, err
	}

	return WeatherData{
		WindSpeed:          data.Current.WindSpeed,
		TemperatureDegrees: data.Current.Temperature,
	}, nil
}

func getDataOpenWeather(city string) (WeatherData, error) {
	resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", city, config.OPEN_WEATHER_API_KEY))
	if err != nil {
        return WeatherData{}, err
    }

	var data OpenWeatherData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
        return WeatherData{}, err
    }

	return WeatherData{
		WindSpeed:          data.Wind.Speed,
        TemperatureDegrees: data.Main.Temp,
    }, nil
}
