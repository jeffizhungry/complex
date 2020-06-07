//go:generate wire gen
// +build wireinject
// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"
)

// DOESNT WORK
// func ProvideClosureCar() Car {
// 	wire.Build(NewCar, NewBody, func() Wheel {
// 		return Wheel{}
// 	})
// 	return Car{}
// }

func ProvideCar() Car {
	wire.Build(NewCar, NewBody, setupWheel)
	return Car{}
}

func setupWheel() Wheel {
	return Wheel{}
}
