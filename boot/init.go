package boot

import (
	"sync"

	durationdata "github.com/BerryHub/models/duration_data"

	"github.com/BerryHub/config"
	"github.com/BerryHub/routes"
	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

var e *echo.Echo

// initEchoRoutes - Si occupa di inizializzare tutte le rotte definite sotto echo
func initEchoRoutes(e *echo.Echo) {

	routes.InitStaticFile(e)
	routes.InitGetRoutes(e)
	routes.InitPostRoutes(e)
}

// InitServer - Si occupa di lanciare l'applicazione con tutte le dovute operazioni iniziali
func InitServer() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		durationdata.InitDurationData()
	}()

	go func() {
		defer wg.Done()

		e = echo.New()
		e.Use(middleware.Recover())
		e.Use(middleware.Gzip())

		initEchoRoutes(e)

	}()

	wg.Wait()

	config := config.GetInstance()

	e.Logger.Fatal(e.Start(config.AppPort))
}
