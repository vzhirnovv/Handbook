package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return 
	}
	if t.Left == nil {
		ch <- t.Value
		if t.Right != nil {
			Walk(t.Right, ch)
		}
		return
	} else {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		node1, node2 := <-ch1, <-ch2
		if (node1 != node2) {
			return false
		}
	}
	return true
}

func main() {
	channel := make(chan int, 10)
	t1 := tree.New(1)
	t2 := tree.New(1)
	go Walk(t1, channel)
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
	go Walk(t2, channel)
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
	var same bool = Same(t1, t2)
	fmt.Println(same)
      }
