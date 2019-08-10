package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	// functions are values too,
	// and they can be passed around just like other values.
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// function values may be used as function arguments and return values
}