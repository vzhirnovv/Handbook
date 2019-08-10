package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices
	s = append(s, 0) // this will push a new value '0' into a nil slice, extending it's length and capacity
	// Note: the slice will refer to a new, bigger array from this point.
	printSlice(s) 

	// The slice grows as needed.
	s = append(s, 1) // push a new value '1' into the slice, extending it's length and capacity to 2
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}