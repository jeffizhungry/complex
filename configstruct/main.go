package main

import "fmt"

type Wheel struct{}

func NewWheel() Wheel {
	return Wheel{}
}

type Body struct{}

func NewBody() Body {
	return Body{}
}

func NewParts() (Wheel, Body) {
	return Wheel{}, Body{}
}

type CarConfig struct {
	w Wheel
	b Body
}

type Car struct {
	w Wheel
	b Body
}

func NewCar(cfg *CarConfig) Car {
	return Car{w: cfg.w, b: cfg.b}
}

func main() {
	fmt.Println("Hello")
}
