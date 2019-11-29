 package domian

type wind struct {
	Speed int `json:"speed"`
	Direction string `json:"deg"`
}

type Weather struct {
	Description []w `json:"weather"`
	Main Main `json:"main"`
	Wind wind `json:"wind"`
}

type w struct {
	desc string `json:"description"`
}

type Main struct {
	Humidity int `json:"humidity"`
	Temp float64 `json:"temp"`
	TempMin float64 `json:"temp_min"`
	TempMax float64 `json:"temp_max"`
}

func (w *Weather) GetTemperature() (float64, float64, float64) {
	return w.Main.Temp, w.Main.TempMin, w.Main.TempMax
}

func (w *Weather) GetCloudiness() string {
	return w.Description
}

func (w *Weather) GetHumidity() int {
	return w.Main.Humidity
}

func (w *Weather) GetWind() (int, string) {
	return w.Wind.Speed,  w.Wind.Direction
}