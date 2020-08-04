package main;

import "fmt";

var c, python, java bool; // variables without initializers take on default values
var var1, var2 int = 1, 2; // initializing variables with a value
// k := 1 shorthand initializing will not work here,
// because every statement outside a function begins with a keyword.

func main() {
	var i int;
	k := 3; // shorthand initialization of a variable
	fmt.Println(c, python, java, i); // default values will be printed here: false false false 0
	fmt.Println(var1, var2); // 1 2
	fmt.Println(k);
}