package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "it didn't work"}
}

// ---result---
// at 2020-06-24 01:13:40.5127801 +0900 KST m=+0.025986101, it didn't work
