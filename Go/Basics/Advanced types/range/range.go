// The range form of the for loop iterates over a slice or map.

// When ranging over a slice, two values are returned for each iteration. 
// The first is the index, and the second is a copy of the element at that index.

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 62, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	dynamicPow := make([]int, 10)
	// use only the index, second value is omitted
	for i := range dynamicPow {
		dynamicPow[i] = 1 << uint(i) // == 2**i
	}
	// use only the index, recommended to omit the second value
	for index, _ := range dynamicPow {
		fmt.Printf("%d\n", index)
	}
	// use return value only
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}