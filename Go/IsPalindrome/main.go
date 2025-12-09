package main

import (
	"fmt"
	"regexp"

	"strings"
)

func main() {
	res := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(res)
}

func isPalindrome(s string) bool {
	result := false
	firstInput := strings.ToLower(s)
	reg, _ := regexp.Compile(`[^\w]`)
	firstInput = reg.ReplaceAllString(firstInput, "")
	secondInput := reverse(firstInput)
	if secondInput == firstInput {
		return true
	}
	return result
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
