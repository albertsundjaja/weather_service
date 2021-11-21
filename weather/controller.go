package weather

import (
	"encoding/json"
	"net/http"
)

func HandleGetWeather(w http.ResponseWriter, r *http.Request) {
    city := r.URL.Query().Get("city")
	if city == "" {
        city = "sydney"
    }

	weatherData, err := GetWeatherData(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
	jsonData, err := json.Marshal(weatherData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}