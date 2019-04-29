package controllers

import (
	"BerryHub/config"
	"BerryHub/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// GetWeatherData - Restituisce le informazioni sul meteo
func GetWeatherData(c echo.Context) error {

	config := config.GetInstance()

	var content interface{}

	weatherData := models.GetWeatherData()

	content, err := weatherData.GetContent()
	if err == nil {
		return c.JSON(200, Response{
			Status:  0,
			Success: true,
			Message: "ok!",
			Content: content,
		})
	}

	body := make(map[string]interface{})
	err = json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		return c.JSON(500, Response{
			Status:  1,
			Success: false,
			Message: "Errore interno",
		})
	}

	req, err := http.NewRequest("GET", config.OpenWeatherMapURL, nil)
	if err != nil {
		return c.JSON(500, Response{
			Status:  3,
			Success: false,
			Message: err.Error(),
		})
	}

	lat := body["lat"].(float64)
	lon := body["lon"].(float64)

	q := req.URL.Query()
	q.Add("appid", config.OpenWeatherMapAPIToken)
	q.Add("units", "metric")
	q.Add("lat", fmt.Sprintf("%f", lat))
	q.Add("lon", fmt.Sprintf("%f", lon))
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return c.JSON(500, Response{
			Status:  4,
			Success: false,
			Message: err.Error(),
		})
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&content); err != nil {
		return c.JSON(500, Response{
			Status:  5,
			Success: false,
			Message: err.Error(),
		})
	}

	weatherData.SetContent(content)

	return c.JSON(200, Response{
		Status:  0,
		Success: true,
		Message: "ok!",
		Content: content,
	})
}
