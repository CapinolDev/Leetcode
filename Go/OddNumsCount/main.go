package main

import "fmt"

func main() {
	res := countOdds(3, 7)
	fmt.Println(res)
}
func countOdds(low int, high int) int {
	totalcount := high - low + 1
	if totalcount%2 == 0 {
		return totalcount / 2

	} else {
		if low%2 == 0 && high%2 == 0 {
			return (totalcount - 1) / 2
		}
	}
	return 1
}

/*
func countOdds(low int, high int) int {
	res := 0
	if low%2 == 0 {
		low++
	}
	for i := low; i <= high; i += 2 {

		res++

	}
	return res
}
*/
