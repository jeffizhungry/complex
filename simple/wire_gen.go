// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func ProvideCar() Car {
	wheel := NewWheel()
	body := NewBody(wheel)
	car := NewCar(body)
	return car
}
