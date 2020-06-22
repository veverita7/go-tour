package main

import "fmt"

func main() {
	var t []int
	fmt.Println(t, len(t), cap(t))
	if t == nil {
		fmt.Println("nil!")
	}

	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])
	fmt.Println("p[:3] ==", p[:3])
	fmt.Println("p[4:] ==", p[4:])

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	a := make([]int, 5)
	b := make([]int, 0, 5)
	c := b[:2]
	d := c[2:5]
	printSlice("a", a)
	printSlice("b", b)
	printSlice("c", c)
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

// ---result---
// [] 0 0
// nil!
// p == [2 3 5 7 11 13]
// p[1:4] == [3 5 7]
// p[:3] == [2 3 5]
// p[4:] == [11 13]
// p[0] == 2
// p[1] == 3
// p[2] == 5
// p[3] == 7
// p[4] == 11
// p[5] == 13
// a len=5 cap=5 [0 0 0 0 0]
// b len=0 cap=5 []
// c len=2 cap=5 [0 0]
// d len=3 cap=3 [0 0 0]
