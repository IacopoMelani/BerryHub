package config

import (
	"testing"

	"github.com/subosito/gotenv"
)

func TestCacheConfigBoot(t *testing.T) {

	gotenv.Load("../.env")

	config := GetCacheConfig()

	if config.StringConnection == "" {
		t.Error("Errore: variabili d'ambiente non caricate correttamente")
	}

}
