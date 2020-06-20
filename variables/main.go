package main

import "fmt"

var x, y, z int
var c, python, java = true, false, "jvm"

func main() {
	fmt.Println(x, y, z)
	fmt.Println(c, python, java)

	jan, feb, mar := "January", "February", "March"
	fmt.Println(jan, feb, mar)
}

// ---result---
// 0 0 0
// true false jvm
// January February March
