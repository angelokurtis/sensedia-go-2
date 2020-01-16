package clima

import "github.com/google/wire"

var Providers = wire.NewSet(NewOpenWeather, wire.Bind(new(Medidor), new(*OpenWeather)))
