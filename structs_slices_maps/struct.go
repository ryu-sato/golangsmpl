package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)

	p := &v
	(*p).Y = 1e3
	p.X = 1e9
	fmt.Println(v)

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p2 = &Vertex{3, 4}
	)
	fmt.Println(v1, p2, v2, v3)
}
