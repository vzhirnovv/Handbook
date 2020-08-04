package main

import (
	"fmt"
	"math"
)

var	x, y int
var f float64
var z uint

func main() {
	x, y = 3, 4
	f = math.Sqrt(float64(x*x + y*y)) // float f = square root of float conversion of x^2 + y^2 = 5
	z = uint(f) // convert float f to unsigned int z
	fmt.Println(x, y, f, z)
}