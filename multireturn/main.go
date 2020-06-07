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

type Car struct {
	w Wheel
	b Body
}

func NewCar(w Wheel, b Body) Car {
	return Car{w: w, b: b}
}

func main() {
	fmt.Println("Hello")
}
