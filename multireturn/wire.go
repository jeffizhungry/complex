// +build wireinject
// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"
)

func ProvideCar() Car {
	wire.Build(NewCar, NewBody, NewWheel)
	return Car{}
}

// DOESNT WORK
// func ProvideCarWithParts() Car {
// 	wire.Build(NewCar, NewParts)
// 	return Car{}
// }
