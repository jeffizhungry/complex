package main

import "fmt"

type Left struct {
	s string
}

type Right struct {
	s string
}

func main() {
	l := Left{"hello"}
	r := Right{}
	fmt.Println(l)
	fmt.Println(r)
	r = l
	fmt.Println(l)
	fmt.Println(r)
}
