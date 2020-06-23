package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(add2(42, 13))
	fmt.Println(add3(1, 10, 111))

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))

	pos, neg := adder(), adder()
	for i := 0; i < 3; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func add2(x int, y int) int {
	return x + y
}

func add3(x, y, z int) int {
	return x + y + z
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// ---result---
// 55
// 122
// 5
// 0 0
// 1 -2
// 3 -6
