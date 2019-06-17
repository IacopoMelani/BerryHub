package durationdata

import (
	"io"
	"net/http"
	"sync"

	"github.com/BerryHub/helpers/request"

	"github.com/BerryHub/config"
)

// WeatherRemoteData - Definisce il tipo di duration data specifico per il meteo
// implementa RemoteData
type WeatherRemoteData struct{}

var weatherData *DurationData
var onceWeather sync.Once

// GetWeatherData - Restituisce l'istanza di DurationData relativo al meteo
func GetWeatherData() *DurationData {
	onceWeather.Do(func() {

		config := config.GetCacheConfig()

		weatherData = new(DurationData)
		weatherData.ddi = WeatherRemoteData{}
		weatherData.sleepMinute = config.OpenWeatherMapTimeToRefresh
		weatherData.Daemon()
	})
	return weatherData
}

// EncodeQueryString - Restituisce la query string encodata per eseguire la richiesta remota
func (w WeatherRemoteData) EncodeQueryString(req *http.Request) {

	config := config.GetCacheConfig()

	q := req.URL.Query()

	q.Add("appid", config.OpenWeatherMapAPIToken)
	q.Add("units", config.OpenWeatherMapUnits)
	q.Add("lat", config.OpenWeatherMapLatitude)
	q.Add("lon", config.OpenWeatherMapLongitude)
	req.URL.RawQuery = q.Encode()
}

// GetBody - Restituisce il body da inserire in una request
func (w WeatherRemoteData) GetBody() io.Reader {
	return nil
}

// GetMethod - Restituisce il metodo della richiesta remota
func (w WeatherRemoteData) GetMethod() string {
	return "GET"
}

// GetURL - Restituisce la url della richiesta remota
func (w WeatherRemoteData) GetURL() string {
	config := config.GetCacheConfig()
	return config.OpenWeatherMapURL
}

// HandlerData - Metodo per il recupero dei dati
func (w WeatherRemoteData) HandlerData() (interface{}, error) {
	content, err := request.GetRemoteData(w)
	return content, err
}
