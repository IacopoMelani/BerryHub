package controllers

import (
	durationdata "github.com/BerryHub/models/duration_data"

	"github.com/labstack/echo"
)

// GetNewsData - Restituisce le infomarzioni sulle news
func GetNewsData(c echo.Context) error {

	newsData := durationdata.GetNewsData()

	return c.JSON(200, Response{
		Status:  0,
		Success: true,
		Message: "ok!",
		Content: newsData.Content,
	})
}

// GetWeatherData - Restituisce le informazioni sul meteo
func GetWeatherData(c echo.Context) error {

	weatherData := durationdata.GetWeatherData()

	return c.JSON(200, Response{
		Status:  0,
		Success: true,
		Message: "ok!",
		Content: weatherData.Content,
	})
}
