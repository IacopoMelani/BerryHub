package durationdata

import (
	"io"
	"net/http"
	"sync"

	"github.com/BerryHub/config"
)

// NewsDurationData - Defisce il tipo di duration data specifico per le news
type NewsDurationData struct{}

var newsData *DurationData
var onceNews sync.Once

// GetNewsData - Restituisce l'istanza di DurationData relativo alle notizie
func GetNewsData() *DurationData {
	onceNews.Do(func() {
		newsData = new(DurationData)
		newsData.dr = new(NewsDurationData)
		newsData.sleepMinute = 60
		newsData.Daemon()
	})
	return newsData
}

func (w NewsDurationData) encodeQueryString(req *http.Request) {

	config := config.GetInstance()

	q := req.URL.Query()
	q.Add("apiKey", config.NewsAPIToken)
	q.Add("language", config.NewsAPILanguage)
	q.Add("country", config.NewsAPICountry)
	q.Add("category", config.NewsAPICategory)
	req.URL.RawQuery = q.Encode()
}

func (w NewsDurationData) getBody() io.Reader {
	return nil
}

func (w NewsDurationData) getMethod() string {
	return "GET"
}

func (w NewsDurationData) getURL() string {
	config := config.GetInstance()
	return config.NewsAPIURL
}
