package localização

import "github.com/google/wire"

var Providers = wire.NewSet(NewIPStack, wire.Bind(new(Localizador), new(*IPStack)))
