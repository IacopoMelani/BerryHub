package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BerryHub/config"
	"github.com/BerryHub/models"
	"github.com/labstack/echo"
)

// GetNewsData - Restituisce le infomarzioni sulle news
func GetNewsData(c echo.Context) error {

	config := config.GetInstance()

	var content interface{}

	newsData := models.GetNewsData()

	content, err := newsData.GetContent()
	if err == nil {
		return c.JSON(200, Response{
			Status:  0,
			Success: true,
			Message: "ok!",
			Content: content,
		})
	}

	req, err := http.NewRequest("GET", config.NewsAPIURL, nil)
	if err != nil {
		return c.JSON(500, Response{
			Status:  3,
			Success: false,
			Message: err.Error(),
		})
	}

	q := req.URL.Query()
	q.Add("apiKey", config.NewsAPIToken)
	q.Add("language", "it")
	q.Add("country", "it")
	q.Add("category", "technology")
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

	newsData.SetContent(content, 60)

	return c.JSON(200, Response{
		Status:  0,
		Success: true,
		Message: "ok!",
		Content: content,
	})

}

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

	weatherData.SetContent(content, 15)

	return c.JSON(200, Response{
		Status:  0,
		Success: true,
		Message: "ok!",
		Content: content,
	})
}
