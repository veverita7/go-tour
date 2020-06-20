package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// if-else 문 밖에서 v 변수는 사용 불가
	return lim
}

// ---result---
// 1.4142135623730951 2i
// 27 >= 20
// 9 20
