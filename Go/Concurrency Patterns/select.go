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
	// Timeout for whole conversation
	timeout := time.After(10 * time.Second)
	// Multiplex channel rewritten via select
	go func() {
		for {
			select {
			case s := <-ch1:
				c <- s
			case s := <-ch2:
				c <- s
			case <-timeout:
				fmt.Println("You talk too much.")
				return
			}
		}
	}()

	return c
}

func main() {
	c := fanIn(speak("lol"), speak("lmao"))

	quit := make(chan bool)
	time.AfterFunc(time.Duration(rand.Intn(10))*time.Second, func() {
		quit <- true
	})

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second): // time.After returns a channel after a specified interval.
			fmt.Println("Too slow, I don't care anymore.")
			return
		case <-quit:
			fmt.Println("Shut up, please.")
			return
		}
	}
}
