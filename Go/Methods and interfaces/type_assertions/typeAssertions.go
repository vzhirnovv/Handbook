package main

import "fmt"

func main() {
	var i interface{} = "hello" // interface value is now string
	fmt.Println(i) // we will see "hello" here

	s := i.(string) // access the interface value's underlying value
	fmt.Println(s)

	s, ok := i.(string) // check if interface value holds a string type value
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok) // f is zero value and ok is false

	//f = i.(float64) // i does not hold a float64, so panic is thrown here
	// this can be remedied by assigning a float64 value to an interface
	i = 0.69

	f, ok = i.(float64)
	fmt.Println(f, ok) // interface now holds value of type float64, so everything is cool
}