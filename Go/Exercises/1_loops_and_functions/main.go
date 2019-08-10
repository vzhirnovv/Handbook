package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := x
	temp := 0.0
	iterationNumber := 1
	for math.Abs(z - temp) >= 0.000000000000001 {
		fmt.Println("iteration", iterationNumber)
		temp = z
		fmt.Printf("temp is %g\n", temp)
		z -= (z*z - x) / (2*z);
		fmt.Printf("z is %g\n", z)
		iterationNumber++
	} 
	return z;
}

func main() {
	fmt.Println("my sqrt: ", sqrt(3))
	fmt.Println("math.Sqrt: ", math.Sqrt(3))
}