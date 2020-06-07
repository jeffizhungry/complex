package main

import (
	"fmt"

	"github.com/k0kubun/pp"
	"go.uber.org/dig"
)

type WheelConfig string

type Wheel struct {
	color string
}

func NewWheel(color WheelConfig) Wheel {
	return Wheel{string(color)}
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
	c := dig.New()
	c.Provide(func() WheelConfig {
		return "blue"
	})
	c.Provide(NewCar)
	c.Provide(NewBody)
	c.Provide(NewWheel)
	err := c.Invoke(func(c Car) {
		pp.Println("this is a car:", c)
	})
	if err != nil {
		fmt.Println(err)
	}
	// if err := dig.Visualize(c, os.Stdout); err != nil {
	// 	fmt.Println(err)
	// }
}
