package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var k = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(k)

	a := make(map[string]int)

	a["Answer"] = 42
	fmt.Println("Answer =", a["Answer"])
	delete(a, "Answer")
	fmt.Println("Answer =", a["Answer"])
	v, ok := a["Answer"]
	fmt.Println("value =", v, "present?", ok)

}
