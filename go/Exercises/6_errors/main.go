package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func sqrt(x float64) (float64, error) {
	if (x < 0) {
		return 0, ErrNegativeSqrt(x)
	}
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
	return z, nil;
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}

