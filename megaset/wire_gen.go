// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
)

// Injectors from wire.go:

func ProvideWheel() Wheel {
	wheel := NewWheel()
	return wheel
}

func ProvideBody() Body {
	wheel := NewWheel()
	body := NewBody(wheel)
	return body
}

func ProvideCar() Car {
	wheel := NewWheel()
	body := NewBody(wheel)
	car := NewCar(body)
	return car
}

// wire.go:

var MegaSet = wire.NewSet(NewCar, NewBody, NewWheel, NewInteger)

func NewInteger() int {
	return 5
}
