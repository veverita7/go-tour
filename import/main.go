package main

// fmt, math 패키지를 임포트
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("now you have %g problems.", math.Nextafter(2, 3))
}

// ---result---
// now you have 2.0000000000000004 problems.
