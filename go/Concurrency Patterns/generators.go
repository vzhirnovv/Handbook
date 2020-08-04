package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator
func speak(message string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", message, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c1 := speak("lol")
	c2 := speak("lmao")

	for i := 0; i < 5; i++ {
		// Lockstep will occur here
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}

	fmt.Println("Done")
}
