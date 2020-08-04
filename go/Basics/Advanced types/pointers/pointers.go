package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i // p now points to memory containing i
	fmt.Println(*p) // print the contents of the memory address p stores
	*p = 21 // change i via pointer
	fmt.Println(i)

	p = &j // p now points to memory containing j
	*p = *p / 37 // divide j via pointer
	fmt.Println(j)

}