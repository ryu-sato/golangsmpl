package main

import (
	"fmt"
)

func fibonacci() func() int {
	a, b, next := 0, 0, 0
	return func() (num int) {
		if (next == 0) {
			num = 0
			a = 0
			b = 1
			next = 1
		} else {
			num = next
			next = a + b
			a = b
			b = next
		}
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
