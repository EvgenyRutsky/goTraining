package main

import (
	"fmt"
	"forecast/domian"
	"os"
	"time"
)

func main () {

	var cityname string

	fmt.Println("Please enter your city:")
	_, err := fmt.Scanln(&cityname)

	if err != nil {
		fmt.Println("Something went wrong during the scanning your input")
		os.Exit(2)
	}

	meteo := domian.Meteorologist{}
	currentWeather := meteo.WeatherForecast(cityname)


	temp, _, _ := currentWeather.GetTemperature()
	wind, direction := currentWeather.GetWind()

	loc := time.FixedZone("UTC", currentWeather.Timezone)
	sunrise := time.Unix(int64(currentWeather.Sys.Sunrise),0)
	sunset := time.Unix(int64(currentWeather.Sys.Sunset),0)

	sunrise = sunrise.In(loc)
	sunset = sunset.In(loc)

	fmt.Printf("Сегодня в городе %v %v, температура воздуха %v°С, ветер %v %v м/с. Влажность воздуха %v. Восход солнца %v, заход солнца %v.\n",
		cityname,
		currentWeather.GetCloudiness(),
		temp,
		direction,
		wind,
		currentWeather.GetHumidity(),
		sunrise.Format("15.04"),
		sunset.Format("15.04"),
	)
}

