package routes

import (
	"github.com/BerryHub/controllers"
	"github.com/labstack/echo"
)

// InitGetRoutes - Dichiara tutte le route GET
func InitGetRoutes(e *echo.Echo) {
	e.GET("user/all", controllers.GetAllUser)
	e.GET("news/data", controllers.GetNewsData)
}

// InitPostRoutes - Dichiara tutte le route POST
func InitPostRoutes(e *echo.Echo) {
	e.POST("weather/data", controllers.GetWeatherData)
}

// InitStaticFile - Dichiara tuttle le route che caricano file statici
func InitStaticFile(e *echo.Echo) {
	e.Static("*", "./dist/")
}
