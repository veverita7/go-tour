package main

import (
	"fmt"
	"math"
	"os"
)

type Abser interface {
	Abs() float64
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	f := MyFloat(-math.Sqrt2)

	var a Abser = &v
	// var a Abser = v
	fmt.Println(a.Abs())
	a = f
	fmt.Println(a.Abs())

	var w Writer = os.Stdout
	fmt.Fprintf(w, "hello writer\n")
}

// ---result---
// 5
// 1.4142135623730951
//  hello writer
