package main

import "fmt"

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// ---result---
// world hello
// 7 10
