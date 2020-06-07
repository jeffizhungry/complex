package main

import "fmt"

type Wheel struct{}

func NewWheel() Wheel {
	return Wheel{}
}

type Body struct {
	w Wheel
}

func NewBody(w Wheel) Body {
	return Body{w}
}

type Car struct {
	b Body
}

func NewCar(b Body) Car {
	return Car{b}
}

func main() {
	fmt.Println("Hello")
}
