package main

import (
	"fmt"
)

func main() {
	input := []int{1, 2, 3}
	res := searchInsert(input, 2)
	fmt.Println(res)
}

func searchInsert(nums []int, target int) int {
	searchPos := len(nums)
	if searchPos%2 == 0 {
		searchPos = (searchPos / 2)
	} else {
		searchPos = int(float32(searchPos/2) + 0.5)
	}
	fmt.Println(searchPos)
	if nums[searchPos] == target {
		return searchPos

	}

	return -10
}
