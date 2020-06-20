package main

import "fmt"

const (
	Pi    = 3.14
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

// ---result---
// Happy 3.14 Day
// Go rules? true
// 21
// 0.2
// 1.2676506002282295e+29
