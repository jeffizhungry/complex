// +build wireinject
// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"
)

var MegaSet = wire.NewSet(NewCar, NewBody, NewWheel, NewInteger)

func ProvideWheel() Wheel {
	wire.Build(MegaSet)
	return Wheel{}
}

func ProvideBody() Body {
	wire.Build(MegaSet)
	return Body{}
}

func ProvideCar() Car {
	wire.Build(MegaSet)
	return Car{}
}

func NewInteger() int {
	return 5
}
