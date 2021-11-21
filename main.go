package main

import (
	"github.com/albertsundjaja/weather_service/router"
	"log"
	"net/http"
)

func main() {
	router := router.CreateRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
