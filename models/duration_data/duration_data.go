package durationdata

import (
	"errors"
	"sync"
	"time"
)

// DDInterface -
type DDInterface interface {
	HandlerData() (interface{}, error)
}

// DurationData - Struct per immagazzinare i dati raccolti con il suo relativo tempo di scadenza dopo il quale è obbligato a ricevere nuovi dati
//in alternativa è possibile definere una fuzione handler da assegnare all'istanza di DurationData, un intervallo di tempo in minuti nel quale l'handler viene richiamato per poi avviare il demone relativo alla stessa istanza
type DurationData struct {
	mu          sync.Mutex
	ddi         DDInterface
	stopSignal  chan bool
	sleepMinute int
	Content     interface{}
	ExpiredAt   time.Time
}

// InitDurationData - Si occupa di avviare tutte le istanze di DurationData
func InitDurationData() {
	GetNewsData()
	GetWeatherData()
}

// getDaemonData - Si occupa di prevelare i dati dall'handler e se non ci sono stati errori lo sostituisce con quello nuovo
func (d *DurationData) getDaemonData() {
	content, err := d.ddi.HandlerData()
	if err == nil {
		d.mu.Lock()
		d.Content = content
		d.mu.Unlock()
	}
}

// Daemon - Si occupa di avviare il demone che aggiorna i dati, esso può essere ucciso richiamando il metodo StopDaemon()
func (d *DurationData) Daemon() {

	d.stopSignal = make(chan bool)

	go func() {
		ticker := time.NewTicker(time.Second * time.Duration(d.sleepMinute))

		d.getDaemonData()

		for {
			select {
			case <-d.stopSignal:
				ticker.Stop()
				return
			case <-ticker.C:
				d.getDaemonData()

			}
		}
	}()
}

// GetContent - Restituisce i dati recuperati nel caso siano presenti e non siano scaduti altrimenti errore
func (d *DurationData) GetContent() (interface{}, error) {

	if d.ExpiredAt.IsZero() || d.Content == nil {
		return nil, errors.New("Dati mancanti")
	}

	diff := d.ExpiredAt.Sub(time.Now())
	if diff.Seconds() <= 0 {
		return nil, errors.New("Data scaduta")
	}
	return d.Content, nil
}

// SetContent - Imposta dei nuovi dati e aggiorando il tempo di scadenza solo se i precedenti non sono più validi, altrimenti non fa niente
func (d *DurationData) SetContent(content interface{}, minutsInterval int) {

	if diff := d.ExpiredAt.Sub(time.Now()); diff.Seconds() > 0 {
		return
	}

	d.mu.Lock()
	d.Content = content
	d.ExpiredAt = time.Now().Add(time.Minute * time.Duration(minutsInterval))
	d.mu.Unlock()
}

// StopDaemon - Si occupa di avvertire il demone di fermarsi
func (d *DurationData) StopDaemon() {
	d.stopSignal <- true
	close(d.stopSignal)
}
