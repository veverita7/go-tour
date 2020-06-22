package main

import "fmt"

type Vertex struct {
	X int
	Y int // == X, Y int
}

func main() {
	s := Vertex{}
	p := Vertex{X: 1}
	q := Vertex{1, 2}
	q.X = 4
	fmt.Println(s, p, q)
}

// ---result---
// {0 0} {1 0} {4 2}
