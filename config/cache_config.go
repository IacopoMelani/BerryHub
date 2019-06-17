package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"sync"
)

// CacheConfig - struttura dove immagazzinare le configurazioni
type CacheConfig struct {
	StringConnection            string
	AppPort                     string
	OpenWeatherMapAPIToken      string
	OpenWeatherMapURL           string
	OpenWeatherMapUnits         string
	OpenWeatherMapLatitude      string
	OpenWeatherMapLongitude     string
	OpenWeatherMapTimeToRefresh int
	NewsAPIToken                string
	NewsAPIURL                  string
	NewsAPILanguage             string
	NewsAPICountry              string
	NewsAPICategory             string
	NewsAPITimeToRefresh        int
}

var arrayEnvMapper = map[string]string{
	"STRING_CONNECTION":              "StringConnection",
	"APP_PORT":                       "AppPort",
	"OPENWEATHERMAP_API_TOKEN":       "OpenWeatherMapAPIToken",
	"OPENWEATHERMAP_URL":             "OpenWeatherMapURL",
	"OPENWEATHERMAP_UNITS":           "OpenWeatherMapUnits",
	"OPENWEATHERMAP_LATITUDE":        "OpenWeatherMapLatitude",
	"OPENWEATHERMAP_LONGITUDE":       "OpenWeatherMapLongitude",
	"OPENWEATHERMAP_TIME_TO_REFRESH": "OpenWeatherMapTimeToRefresh",
	"NEWS_API_TOKEN":                 "NewsAPIToken",
	"NEWS_API_URL":                   "NewsAPIURL",
	"NEWS_API_LANGUAGE":              "NewsAPILanguage",
	"NEWS_API_COUNTRY":               "NewsAPICountry",
	"NEWS_API_CATEGORY":              "NewsAPICategory",
	"NEWS_API_TIME_TO_REFRESH":       "NewsAPITimeToRefresh",
}

var cacheConfig *CacheConfig
var once sync.Once

// GetCacheConfig - restituisce l'unica istanza della struttura contenente le configurazioni
func GetCacheConfig() *CacheConfig {
	once.Do(func() {
		cacheConfig = &CacheConfig{}
		cacheConfig.loadEnvConfig()

	})
	return cacheConfig
}

// loadEnvConfig - si occupa di caricare tutte le configurazioni dell'env nella struttura di configurazione
func (c *CacheConfig) loadEnvConfig() {
	for envName, StructName := range arrayEnvMapper {
		err := c.setField(StructName, os.Getenv(envName))
		if err != nil {
			panic("Errore caricamento configurazione")
		}
	}
}

// setField - si occupa di impostare  attrun campo averso la reflection, c รจ necessario sia un puntatore a una struttura
func (c *CacheConfig) setField(name string, value string) error {

	rv := reflect.ValueOf(c)

	// Controllo se pointer a una struct
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return errors.New("c must be pointer to struct")
	}

	// Prelevo i campi della struct
	rv = rv.Elem()

	// Controllo che il campo esista
	fv := rv.FieldByName(name)
	if !fv.IsValid() {
		return fmt.Errorf("not a field name: %s", name)
	}

	if !fv.CanSet() {
		return fmt.Errorf("cannot set field %s", name)
	}

	// Controllo tipo stringa
	if fv.Kind() == reflect.String {
		fv.SetString(value)
		return nil
	}
	if fv.Kind() == reflect.Int {
		content, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return errors.New("Invalid value for int")
		}
		fv.SetInt(content)
		return nil
	}

	return errors.New("Invalid type for " + name)
}
