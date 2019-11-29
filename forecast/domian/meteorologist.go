package domian

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	scheme = "http"
	host = "api.openweathermap.org"
	path = "data/2.5/weather"
	langParam = "ru"
	unitsParam = "metric"
	accessKey = "2c19a8c670afc70f2ae7a81f229fce3d"
)

type Meteorologist struct {

}

func (m Meteorologist) WeatherForecast (city string) Weather {

	u := urlBuilder(city)

	fmt.Println(u)

	resp, err := http.Get(u)

	if err != nil {
		fmt.Printf("An error occurred while requesting to the url %v\n", u)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	if err != nil {
		fmt.Println("An error occurred while reading response body")
	}

	weather := Weather{}

	err = json.Unmarshal(body, &weather)

	if err != nil {
		fmt.Println("An error occurred while parsing response", err)
	}
	
	return weather
}

func urlBuilder (city string) string {

	u := &url.URL{
		Scheme:     scheme,
		Host:       host,
		Path:       path,
	}

	q := u.Query()
	q.Set("q", city)
	q.Add("lang", langParam)
	q.Add("units", unitsParam)
	q.Add("appid", accessKey)
	u.RawQuery = q.Encode()

	return u.String()
}