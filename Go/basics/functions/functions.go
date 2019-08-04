package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func addOmitted(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(2, 3))
}
