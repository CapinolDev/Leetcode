package main

import (
	"fmt"
	"math"
)

func main() {
	res := maximumStrongPairXor([]int{1, 2, 3, 4, 5})
	fmt.Println(res)
}
func maximumStrongPairXor(nums []int) int {
	maxXor := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if math.Abs(float64(nums[i]-nums[j])) <= float64(min(nums[i], nums[j])) {
				maxXor = max((nums[i] ^ nums[j]), maxXor)
			}
		}
	}
	return maxXor
}
