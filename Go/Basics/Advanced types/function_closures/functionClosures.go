package main

import "fmt"

// the adder function return a closure
// each closure is bound to its own sum variable
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder() // even though pos and neg are assign the same closure, they will have different sum variables assigned to them
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), // positive sum
			neg(-2*i), // negative sum
		)
	}
}
