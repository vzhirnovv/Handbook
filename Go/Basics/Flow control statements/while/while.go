package main

import "fmt"

func main() {
	sum := 1
	// for is Go's while 
	for sum < 1000 { // <==> while sum < 1000
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)
}