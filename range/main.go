package main

import "fmt"

var pow = []int{1, 2, 4, 8}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow = make([]int, 4)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// ---result---
// 2**0 = 1
// 2**1 = 2
// 2**2 = 4
// 2**3 = 8
// 1
// 2
// 4
// 8