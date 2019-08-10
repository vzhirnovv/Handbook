package main

import "fmt"

func main() {
	sum := 0;
	// no parentheses and braces are ALWAYS required
	for i := 0; i < 10; i++ {
		sum += i;
	}
	fmt.Println(sum)
	
	// init and post statements are optional
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}