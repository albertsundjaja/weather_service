package config

import (
	"os"
	"github.com/joho/godotenv"
)

var OPEN_WEATHER_API_KEY = ""
var WEATHER_STACK_API_KEY = ""

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	OPEN_WEATHER_API_KEY = os.Getenv("OPEN_WEATHER_API_KEY")
    WEATHER_STACK_API_KEY = os.Getenv("WEATHER_STACK_API_KEY")
}