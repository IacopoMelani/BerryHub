package db

import (
	"database/sql"

	"github.com/BerryHub/config"
	// Carica il driver mysql per la connessione al db
	_ "github.com/go-sql-driver/mysql"

	"log"
	"sync"
)

var db *sql.DB

var once sync.Once

// GetConnection - restituisce un'istanza di connessione al database
func GetConnection() *sql.DB {

	once.Do(func() {

		config := config.GetCacheConfig()
		conn, err := sql.Open("mysql", config.StringConnection)
		if err != nil {
			log.Panic(err.Error())
		}

		if err := conn.Ping(); err != nil {
			log.Panic(err.Error())
		}

		db = conn
	})

	return db
}
