package durationdata

import (
	"io"
	"net/http"
	"sync"

	"github.com/BerryHub/config"
)

// NewsRemoteData - Defisce il tipo di remote data specifico per le news
// implementa RemoteData
type NewsRemoteData struct{}

var newsData *DurationData
var onceNews sync.Once

// GetNewsData - Restituisce l'istanza di DurationData relativo alle notizie
func GetNewsData() *DurationData {
	onceNews.Do(func() {
		newsData = new(DurationData)
		newsData.rd = NewsRemoteData{}
		newsData.sleepMinute = 60
		newsData.Daemon()
	})
	return newsData
}

// EncodeQueryString - Restituisce la query string encodata per eseguire la richiesta remota
func (w NewsRemoteData) EncodeQueryString(req *http.Request) {

	config := config.GetCacheConfig()

	q := req.URL.Query()
	q.Add("apiKey", config.NewsAPIToken)
	q.Add("language", config.NewsAPILanguage)
	q.Add("country", config.NewsAPICountry)
	q.Add("category", config.NewsAPICategory)
	req.URL.RawQuery = q.Encode()
}

// GetBody - Restituisce il body da inserire in una request
func (w NewsRemoteData) GetBody() io.Reader {
	return nil
}

// GetMethod - Restituisce il metodo della richiesta remota
func (w NewsRemoteData) GetMethod() string {
	return "GET"
}

// GetURL - Restituisce la url della richiesta remota
func (w NewsRemoteData) GetURL() string {
	config := config.GetCacheConfig()
	return config.NewsAPIURL
}
