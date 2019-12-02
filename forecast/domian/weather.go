 package domian

 import "math"

 type Weather struct {
	WeatherDescription []WeatherDescription `json:"weather"`
	Main Main `json:"main"`
	Wind Wind `json:"wind"`
	Sys Sys `json:"sys"`
	Timezone int `json:"timezone"`
}

type WeatherDescription struct {
	Description string `json:"description"`
}

type Main struct {
	Temp float64 `json:"temp"`
	Humidity int `json:"humidity"`
	TempMin float64 `json:"temp_min"`
	TempMax float64 `json:"temp_max"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Direction int `json:"deg"`
}

type Sys struct {
	Sunrise int `json:"sunrise"`
	Sunset int `json:"sunset"`
}

func (w *Weather) GetTemperature() (float64, float64, float64) {
	return w.Main.Temp, w.Main.TempMin, w.Main.TempMax
}

func (w *Weather) GetCloudiness() string {

	for _, v := range w.WeatherDescription {
		return v.Description
	}

	return ""
}

func (w *Weather) GetHumidity() int {
	return w.Main.Humidity
}

func (w *Weather) GetWind() (float64, string) {
	direction := directionToText(w.Wind.Direction)
	return w.Wind.Speed, direction
}

func directionToText(deg int) string {
	v := math.Floor((float64(deg)/45) + 0.5)
	m := map[float64] string {
		0 : "северный",
		1 : "северо-восточный",
		2 : "восточный",
		3 : "юго-восточный",
		4 : "южный",
		5 : "юго-западный",
		6 : "западный",
		7 : "северо-западный",
		8 : "северный",
	}

	return m[v]
}