package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// struct literals 
var (
	sv1 = Vertex{1, 2} // has type Vertex
	sv2 = Vertex{X: 1} // Y:0 is implicit
	sv3 = Vertex{} // X:0, Y:0
	sp = &Vertex{1, 2} // has type *Vertex, returns a pointer to the struct value
)

func main() {
	fmt.Println(Vertex{1, 2})
	
	// struct field access
	v1 := Vertex{1, 2}
	v1.X = 4
	fmt.Println(v1)

	// pointers to structs	
	v2 := Vertex{1, 2}
	p := &v2 // p now points to a memory address that contains the strucs
	p.X = 1e9 // notice that dereference is not required here ((*p).X) is optional, just p.X can be used)
	fmt.Println(v2)

	// struct literals
	fmt.Println(sv1, sp, sv2, sv3)
}