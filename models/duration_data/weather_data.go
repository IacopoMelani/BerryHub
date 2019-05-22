package durationdata

import (
	"io"
	"net/http"
	"sync"

	"github.com/BerryHub/config"
)

// WeatherDurationData - Definisce il tipo di duration data specifico per il meteo
type WeatherDurationData struct{}

var weatherData *DurationData
var onceWeather sync.Once

// GetWeatherData - Restituisce l'istanza di DurationData relativo al meteo
func GetWeatherData() *DurationData {
	onceWeather.Do(func() {
		weatherData = new(DurationData)
		weatherData.sleepMinute = 15
		weatherData.dr = new(WeatherDurationData)
		weatherData.Daemon()
	})
	return weatherData
}

func (w WeatherDurationData) encodeQueryString(req *http.Request) {

	config := config.GetInstance()

	q := req.URL.Query()

	q.Add("appid", config.OpenWeatherMapAPIToken)
	q.Add("units", config.OpenWeatherMapUnits)
	q.Add("lat", config.OpenWeatherMapLatitude)
	q.Add("lon", config.OpenWeatherMapLongitude)
	req.URL.RawQuery = q.Encode()
}

func (w WeatherDurationData) getBody() io.Reader {
	return nil
}

func (w WeatherDurationData) getMethod() string {
	return "GET"
}

func (w WeatherDurationData) getURL() string {
	config := config.GetInstance()
	return config.OpenWeatherMapURL
}
