package boot

import (
	"sync"

	// dsf
	_ "expvar"
	// sd

	// SDFSD
	"net/http"
	_ "net/http/pprof"

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

	wg.Add(3)

	go func() {
		defer wg.Done()
		config.GetInstance()
	}()

	go func() {
		defer wg.Done()
		//durationdata.InitDurationData()
	}()

	go func() {
		defer wg.Done()

		e = echo.New()
		e.Use(middleware.Recover())
		e.Use(middleware.Gzip())

		e.GET("/debug/pprof/*", echo.WrapHandler(http.DefaultServeMux))

		initEchoRoutes(e)

	}()

	wg.Wait()

	config := config.GetInstance()

	e.Logger.Fatal(e.Start(config.AppPort))
}
