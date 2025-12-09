package main

import "fmt"

func main() {
	result := mySqrt(1)
	fmt.Println(result)
}

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	low := 1
	high := x
	ans := 0

	for low <= high {
		mid := low + (high-low)/2
		square := mid * mid

		if square == x {
			return mid
		} else if square < x {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}
