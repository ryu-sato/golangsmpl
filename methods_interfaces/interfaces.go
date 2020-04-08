package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	
	a = f   // a MyFloat implements Abser
	a = &v  // a *Vertex implements Abser
	
	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v
	
	fmt.Println(a.Abs())
	
	var i I2 = T{"hello"}
	i.M2()
	
	var j I
	
	j = &T{"Hello"}
	describe(j)
	j.M()
	
	j = F(math.Pi)
	describe(j)
	j.M()
	
	var k I
	var t *T
	k = t
	describe(k)
	k.M()
	
	k = &T{"hellohello"}
	describe(k)
	k.M()
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type I2 interface {
	M2()
}

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M2() {
	fmt.Println(t.S)
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
