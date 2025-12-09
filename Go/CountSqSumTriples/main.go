package main

import (
	"fmt"
)

func main() {
	res := countTriples(5)

	fmt.Println(res)
}

func countTriples(n int) int {
	var a int
	var b int
	var c int
	result := 0

	for a = 2; a <= n; a++ {
		for b = 2; b <= n; b++ {
			for c = 2; c <= n; c++ {
				if (a*a)+(b*b) == (c * c) {
					fmt.Println(a, b, c)
					result++
				}
			}
		}
	}

	return result
}
