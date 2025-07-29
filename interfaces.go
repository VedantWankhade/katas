package katas

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	h int
	b int
}

func (r *rectangle) area() float64 {
	return float64(r.h * r.b)
}

type circle struct {
	r int
}

func (c circle) area() float64 {
	return math.Pi * float64(c.r*c.r)
}

func interfaces() {
	r := rectangle{2, 4}
	c := circle{3}
	printArea(&r) // works
	// printArea(r) // this gives error
	printArea(&c) // works
	printArea(c)  // also works

}

func printArea(s shape) {
	fmt.Println(s.area())
}
