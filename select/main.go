package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)

	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}

// ---result---
// 0
// 1
// 1
// 2
// 3
// quit
//     .
//     .
// tick.
//     .
//     .
// tick.
//     .
//     .
// tick.
//     .
//     .
// tick.
//     .
//     .
// tick.
// boom!
