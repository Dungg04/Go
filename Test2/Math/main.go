package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)

	c := math.Max(a, b)
	fmt.Println(c)

	d := math.Mod(a, b)
	fmt.Println(d)
}