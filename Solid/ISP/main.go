package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float32
}

type object interface {
	shape 
	volume() float32
}

type square struct {
	sideLen float32
}

func (s square) area() float32 {
	return s.sideLen*s.sideLen
}

type cube struct {
	square
}

func (c cube) volume() float32 {
	return math.Pow(c.sideLen, 3)
}

type calculator struct {
	total float32
}

func (c *calculator) sumAreas(shapes ...shape) float32 {
	var sum float32

	for _, shape := range shapes {
		sum += shape.area()
	}

	return sum
}

func (c *calculator) sumVolumes(shapes ...object) float32 {
	var sum float32

	for _, shape := range shapes {
		sum += shape.area() + shape.volume()
	}

	return sum
} 

func main() {
	s := square{sideLen: 2}
	c := cube{square{sideLen: 3}}

	calculator := calculator{}

	fmt.Println(calculator.sumAreas(s))
	fmt.Println(calculator.sumVolumes(c))
}