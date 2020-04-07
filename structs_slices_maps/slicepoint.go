package main

import (
	"fmt"
)

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	/* slice literals */
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	
	/* Slice bound defaults */
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	
	t := []int{2, 3, 5, 7, 11, 13}
	t = t[1:4]
	fmt.Println(t)
	t = t[:2]
	fmt.Println(t)
	t = t[1:]
	fmt.Println(t)
	
	/* slice length and capacity */
	u := []int{2, 3, 5, 7, 11, 13}
	printSlice(u)
	
	u = u[:0]
	printSlice(u)
	
	u = u[:4]
	printSlice(u)
	
	u = u[1:]
	printSlice(u)
	
	var v []int
	printSlice(v)
	if v == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)	
}
