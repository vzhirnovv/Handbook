package main

import "fmt"

func printSlice(s []int) {
	fmt.Println(s)
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s))
}

func main() {
	fmt.Println("Array basics\n")
	// variable declaration
	var a[2] string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// shorthand declaration
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes, "\n")

	fmt.Println("Array slicing\n")
	var s []int = primes[1:4] // everything between elements 1 and 4, including element 1 and excluding 4
	fmt.Println(s, "\n")

	// changing slices
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	namesSlice1 := names[0:2]
	namesSlice2 := names[1:3]
	fmt.Println(namesSlice1, namesSlice2)

	namesSlice2[0] = "XXX"
	fmt.Println(namesSlice1, namesSlice2)
	fmt.Println(names, "\n")

	// slice literals
	fmt.Println("Slice literals\n")
	s1 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s1)

	s2 := []bool{true, false, true, true, false, true}
	fmt.Println(s2)

	s3 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s3, "\n")

	// slice length and capacity
	fmt.Println("Slice length and capacity\n")
	s4 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s4)
	s4 = s4[:0]
	printSlice(s4)
	s4 = s4[:4]
	printSlice(s4)
	s4 = s4[2:]
	printSlice(s4)
	fmt.Println()
	// nil slices
	fmt.Println("Nil slices\n")
	var nilSlice []int
	printSlice(nilSlice)
}