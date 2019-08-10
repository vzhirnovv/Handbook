package main

import "fmt"

func promise1() {
	fmt.Println("this is the first function to be executed, second function will be executed only after this one.")
}

func promise2() {
	defer fmt.Println("this is the second function to be executed.")

	promise1()
}

func main() {
	promise2()

	fmt.Println("The following output is an example of defer stacking.")


	fmt.Println("counting...")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done.")
}