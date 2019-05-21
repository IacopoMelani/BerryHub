package durationdata

import "sync"

// WeatherDurationData - Definisce il tipo di duration data specifico per il meteo
type WeatherDurationData DurationData

var weatherData *DurationData
var onceWeather sync.Once

// GetWeatherData - Restituisce l'istanza di DurationData relativo al meteo
func GetWeatherData() *DurationData {
	onceWeather.Do(func() {
		weatherData = new(DurationData)
	})
	return weatherData
}
