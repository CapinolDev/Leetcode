package main

import (
	"fmt"
	"math"
)

func main() {
	res := isPowerOfTwo(1)
	fmt.Println(res)
}

func isPowerOfTwo(n int) bool {
	for i := 0; i < n; i++ {
		if int(math.Pow(2, float64(i))) == n {
			return true
		}
	}
	return false
}
