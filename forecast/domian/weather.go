 package domian

type temperature struct {
	Temp float64 `json:"main.temp"`
	TempMin float64 `json:"main.temp_min"`
	TempMax float64 `json:"main.temp_max"`
}

type wind struct {
	Speed int `json:"wind.speed"`
	Gust int
	Direction string `json:"wind.deg"`
}

type Weather struct {
	Temperature temperature
	Description string `json:"weather.description"`
	Humidity int `json:"main.humidity"`
	Wind wind
}

func (w *Weather) GetTemperature() (float64, float64, float64) {
	return w.Temperature.Temp, w.Temperature.TempMin, w.Temperature.TempMax
}

func (w *Weather) GetCloudiness() string {
	return w.Description
}

func (w *Weather) GetHumidity() int {
	return w.Humidity
}

func (w *Weather) GetWind() (int, int, string) {
	return w.Wind.Speed, w.Wind.Gust, w.Wind.Direction
}