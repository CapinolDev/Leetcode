package main

import "fmt"

func main() {
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(heights)
	fmt.Println(result)
}

func maxArea(height []int) int {
	maxHeight := 0
	h1 := 0
	h2 := len(height) - 1

	for h1 < h2 {
		maxHeight = max(maxHeight, min(height[h1], height[h2])*(h2-h1))
		if height[h1] > height[h2] {
			h2--
		} else {
			h1++
		}

	}
	return maxHeight
}
