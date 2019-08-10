package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

// A method can be declared on non-struct types as well.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} 
	return float64(f)
}

func (v Vertex) Abs1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Abs1 and Abs2 work the same way

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs1())
	fmt.Println(Abs2(v))

	f := MyFloat(math.Sqrt2) // convert math.Sqrt2 to myFloat (basically float64 to float64)
	fmt.Println(f.Abs())
}