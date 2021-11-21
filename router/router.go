package router

import (
	"github.com/albertsundjaja/weather_service/weather"
	"github.com/go-chi/chi/v5"
)

func CreateRouter() *chi.Mux {
	v1 := chi.NewRouter()
	v1.Get("/weather", weather.HandleGetWeather)
	
	r := chi.NewRouter()
	r.Mount("/v1", v1)

	return r
}
