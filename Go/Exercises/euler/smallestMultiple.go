package main

import (
	"fmt"
)

func isDivisible(number int) bool {
	for i := 1; i <= 20; i++ {
		if number%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	number := 20
	for {
		if isDivisible(number) {
			fmt.Println(number)
			break
		}
		number += 20
	}
}
