package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
// This function has a pointer reciever
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
// This function DOES NOT have a pointer reciever
func ScaleFunc (v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// This function has a value reciever
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
// This function DOES NOT have a value reciever
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10) // Scale can take a value as the reciever
	ScaleFunc(&v, 10) // ScaleFunc can only take an address
	fmt.Println(v)

	p := &Vertex{4, 3} // p contains an address to the Vertex struct
	p.Scale(10) // Scale can take an adress as the reciever
	ScaleFunc(p, 10)
	fmt.Println(*p)
	// v.Scale(5) == &v.Scale(5) because Scale has a pointer reciever.
	// the same principle takes place in the reverse direction
	v = Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p = &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p)) // AbsFunc has to take a pointer here

	fmt.Println(v, p)
}