package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	// make a slice(dynamic array) with 5 zeroes and capacity of 5
	a := make([]int, 5)
	printSlice("a", a)

	// make a slice with a length of 0 and capacity of 5
	b := make([]int, 0, 5)
	printSlice("b", b)

	// get a slice of b with a length of 2 and capacity of 5 (contains zeroes due to zero values)
	c := b[:2]
	printSlice("c", c)

	// get a slice of b with a length of 3 and capacity of 3 (2nd, 3rd and 4th indices of original slice)
	d := c[2:5]
	printSlice("d", d)
}