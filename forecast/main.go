package main

import (
	"fmt"
	"forecast/domian"
)

func main () {
	meteo := domian.Meteorologist{}

	currentW := meteo.WeatherForecast("Mahilyow")

	fmt.Println(currentW.GetTemperature())
}
