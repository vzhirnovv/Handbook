package main

import (
	"fmt"
)

func main() {
	var sum, prevNum, nextNum, temp int = 0, 0, 1, 0

	for nextNum < 4000000 {
		temp = prevNum + nextNum
		if temp%2 == 0 {
			sum += temp
		}
		prevNum = nextNum
		nextNum = temp
	}
	fmt.Println(sum)
}
