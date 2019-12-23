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

// Multiplex function
func fanIn(ch1, ch2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-ch1
		}
	}()

	go func() {
		for {
			c <- <-ch2
		}
	}()

	return c
}

func main() {
	c := fanIn(speak("lol"), speak("lmao"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Done")
}
