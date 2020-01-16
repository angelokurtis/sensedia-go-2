//+build wireinject

package main

import (
	"github.com/angelokurtis/sensedia-go-2/pkg/clima"
	"github.com/angelokurtis/sensedia-go-2/pkg/localização"
	"github.com/google/wire"
)

func Initialize() clima.Medidor { // Wire's Injector
	wire.Build(clima.Providers, localização.Providers)
	return nil
}
