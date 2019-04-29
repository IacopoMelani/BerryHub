package models

import (
	"errors"
	"sync"
	"time"
)

// WeatherData - Struct per immagazzinare i dati raccolti con il suo relativo tempo di scadenza dopo il quale è obbligato a ricevere nuovi dati
type WeatherData struct {
	Content   interface{}
	ExpiredAt time.Time
}

var weatherData *WeatherData
var onceWeather sync.Once

// GetWeatherData - Restituisce l'istanza di WeatherData
func GetWeatherData() *WeatherData {
	onceWeather.Do(func() {
		weatherData = new(WeatherData)
	})
	return weatherData
}

// GetContent - Restituisce i dati recuperati nel caso siano presenti e non siano scaduti altrimenti errore
func (w *WeatherData) GetContent() (interface{}, error) {

	if w.ExpiredAt.IsZero() || w.Content == nil {
		return nil, errors.New("Dati mancanti")
	}

	diff := w.ExpiredAt.Sub(time.Now())
	if diff.Seconds() <= 0 {
		return nil, errors.New("Data scaduta")
	}
	return w.Content, nil
}

// SetContent - Imposta dei nuovi dati e aggiorando il tempo di scadenza solo se i precedenti non sono più validi, altrimenti non fa niente
func (w *WeatherData) SetContent(content interface{}) {

	if w.ExpiredAt.IsZero() {
		w.Content = content
		w.ExpiredAt = time.Now().Add(time.Minute * 15)
		return
	}

	if diff := w.ExpiredAt.Sub(time.Now()); diff.Seconds() > 0 {
		return
	}

	w.Content = content
	w.ExpiredAt = time.Now().Add(time.Minute * 15)
}
