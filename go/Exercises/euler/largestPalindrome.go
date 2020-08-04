package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(number string) bool {
	for i := 0; i < len(number)/2; i++ {
		if number[i] != number[len(number)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	maxPalindrome := -1
	for i := 9999; i > 999; i-- {
		for j := 9999; j > 999; j-- {
			if isPalindrome(strconv.Itoa(i*j)) && i*j > maxPalindrome {
				maxPalindrome = i * j
			}
		}
	}
	fmt.Println(maxPalindrome)
}
