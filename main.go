package main

import (
	"github.com/angelokurtis/sensedia-go-2/internal/log"
	móduloDeClima "github.com/angelokurtis/sensedia-go-2/pkg/clima"
)

func main() {
	medidor := móduloDeClima.NewOpenWeather() // TODO: use Google's wire initializer to provide the implementations
	clima, err := medidor.AferirMinhaLocalização()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("o clima atual é de %.2f°C e %s\n", clima.Temperatura, clima.Descrição)
}
