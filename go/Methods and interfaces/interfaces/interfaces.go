package main 

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
// MyFloat type now implements Abser
type Vertex struct {
	X, Y float64
}
// *Vertex type now implements Abser
func (v *Vertex ) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a1 Abser
	var a2 Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a1 = f // a MyFloat implements Abser
	a2 = &v // a *Vertex implements Abser

	//a = v THIS WILL NOT WORK, becuase Vertex does not implement Abser, 
	//		only *Vertex does.
	fmt.Println(a1, a2)
	fmt.Println(a1.Abs(), a2.Abs())
}