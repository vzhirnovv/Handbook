package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var prev, next int
	next = 1
	return func() int {
		current := prev
		prev, next = next, prev + next
		return current
	} 
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}