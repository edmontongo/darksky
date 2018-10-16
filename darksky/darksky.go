package darksky

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Location struct {
	Lat, Long float64
}

type DarkSky struct {
	Host, SecretKey string
}

type Weather struct {
	Currently Currently `json:"currently"`
}

type Currently struct {
	Summary     string  `json:"summary"`
	Temperature float64 `json:"temperature"`
}

var (
	ErrUnauthorized = errors.New("permission denied")
	ErrUnknown      = errors.New("unknown")
)

func NewDarkSky(host, secretKey string) *DarkSky {
	return &DarkSky{Host: host, SecretKey: secretKey}
}

func (ds *DarkSky) Forecast(l Location) (*Weather, error) {
	url := fmt.Sprintf("%v/forecast/%v/%v,%v", ds.Host, ds.SecretKey, l.Lat, l.Long)

	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		switch r.StatusCode {
		case http.StatusUnauthorized:
			return nil, ErrUnauthorized
		default:
			return nil, ErrUnknown
		}
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	w := &Weather{}
	err = json.Unmarshal(b, &w)
	return w, err
}
