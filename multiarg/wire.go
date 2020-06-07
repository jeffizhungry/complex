// +build wireinject
// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"
)

func ProvideCar(c Color, s Style) Car {
	wire.Build(NewCar, NewBody, NewWheel)
	return Car{}
}
