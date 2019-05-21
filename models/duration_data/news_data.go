package durationdata

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/BerryHub/config"
)

// NewsDurationData - Defisce il tipo di duration data specifico per le news
type NewsDurationData DurationData

var newsData *DurationData
var onceNews sync.Once

// getNews - Si occupa di chiamare le api remote di newsAPi.org per farsi restituire le notizie
func getNews() (interface{}, error) {

	config := config.GetInstance()

	req, err := http.NewRequest("GET", config.NewsAPIURL, nil)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	defer res.Body.Close()

	var content interface{}

	if err := json.NewDecoder(res.Body).Decode(&content); err != nil {
		return nil, err
	}

	return content, nil
}

// GetNewsData - Restituisce l'istanza di DurationData relativo alle notizie
func GetNewsData() *DurationData {
	onceNews.Do(func() {
		newsData = new(DurationData)
		newsData.rechargeData = getNews
		newsData.sleepMinute = 60
		newsData.Daemon()
	})
	return newsData
}
