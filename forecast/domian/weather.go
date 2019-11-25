package domian

type temperature struct {
	temp float64
	tempMin float64
	tempMax float64
}

type wind struct {
	speed int
	gust int
	direction string
}

type Weather struct {
	temperature temperature
	description string
	humidity int
	wind wind
}

func (w Weather) GetTemperature() (float64, float64, float64) {
	return w.temperature.temp, w.temperature.tempMin, w.temperature.tempMax
}

func (w Weather) GetCloudiness() string {
	return w.description
}

func (w Weather) GetHumidity() int {
	return w.humidity
}

func (w Weather) GetWind() (int, int, string) {
	return w.wind.speed, w.wind.gust, w.wind.direction
}