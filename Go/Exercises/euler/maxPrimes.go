package main

import (
	"fmt"
)

func main() {
	var n int64 = 600851475143
	primes := []int64{}

	for n % 2 == 0 {
		primes = append(primes, 2)
		n = n/2
	}

	var i int64
	for i = 3; i*i <= n; i+=2 {
		for n % i == 0 {
			primes = append(primes, i)
			n = n/i
		}
	}

	if n > 2 {
		primes = append(primes, n)
	}

	fmt.Println(primes)
}