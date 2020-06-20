package main

import "fmt"

func main() {
	fmt.Println(add2(42, 13))
	fmt.Println(add3(1, 10, 111))
}

func add2(x int, y int) int {
	return x + y
}

func add3(x, y, z int) int {
	return x + y + z
}

// ---result---
// 55
// 122
