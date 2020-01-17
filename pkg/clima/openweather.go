package clima

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/angelokurtis/sensedia-go-2/internal/log"

	"github.com/angelokurtis/sensedia-go-2/pkg/localização"
)

type OpenWeather struct {
	location localização.Localizador
}

func NewOpenWeather() *OpenWeather {
	l := localização.NewIPStack()
	return &OpenWeather{location: l}
}

func (o *OpenWeather) AferirMinhaLocalização() (*Clima, error) {
	log.Info("verificando clima na localidade atual")
	l, err := o.location.OndeEstou()
	if err != nil {
		return nil, err
	}
	return o.AferirCoordenadas(l.Latitude, l.Longitude)
}

func (o *OpenWeather) AferirCoordenadas(lat float64, lon float64) (*Clima, error) {
	log.Infof("verificando clima atual nas coordenadas de latitude %.2f e longitude %.2f", lat, lon)

	ak := os.Getenv("OPEN_WEATHER_API_ACCESS_KEY")
	endpoint := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&APPID=%s&units=metric&lang=pt", lat, lon, ak)
	r, err := http.Get(endpoint)
	if err != nil {
		return nil, err // TODO: add context do the HTTP error
	}
	if r.StatusCode == http.StatusOK {
		b := r.Body
		defer b.Close()
		w := &OpenWeatherResponse{}
		err = json.NewDecoder(b).Decode(w)
		if err != nil {
			return nil, err // TODO: add context to the unmarshalling error
		}
		weathers := w.Weather
		if len(weathers) < 1 {
			return nil, errors.New("weathers is empty")
		}
		weather := weathers[0]
		log.Info("clima atual encontrado")
		return &Clima{
			Descrição:   weather.Description,
			Temperatura: w.Main.Temp,
		}, nil
	}
	return nil, errors.New("OpenWeather answered unexpectedly")
}

type OpenWeatherResponse struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}
