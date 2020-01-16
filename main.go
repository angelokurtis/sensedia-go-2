package main

import (
	"github.com/angelokurtis/sensedia-go-2/internal/log"
)

func main() {
	medidor := Initialize()
	clima, err := medidor.AferirMinhaLocalização()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("o clima atual é de %.2f°C e %s\n", clima.Temperatura, clima.Descrição)
}
