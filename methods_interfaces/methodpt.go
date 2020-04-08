package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}

	ScaleFunc(&v, 10)
	fmt.Println(v)
	fmt.Println(AbsFunc(v))
	
	v2 := Vertex{3, 4}
	v2.Scale(10)
	fmt.Println(v2)
	fmt.Println(AbsFunc(v2))
	p2 := &v2
	p2.Scale(10)
	fmt.Println(p2)
	fmt.Println(AbsFunc(*p2))

	v3 := Vertex{3, 4}
	fmt.Println(v3)
	fmt.Println(AbsFunc(v3))
	fmt.Println(v3.Abs())
	p3 := &Vertex{4, 3}
	fmt.Println(p3)
	fmt.Println(AbsFunc(*p3))
	fmt.Println(p3.Abs())
	
	v4 := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v4, v4.Abs())
	v4.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v4, v4.Abs())
}
