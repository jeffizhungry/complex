package main

import "fmt"

type Color string
type Style string

type Wheel struct {
	color string
	style string
}

func NewWheel(color Color, style Style) Wheel {
	return Wheel{string(color), string(style)}
}

func NewWheelBad(color, style string) Wheel {
	return Wheel{color, style}
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
