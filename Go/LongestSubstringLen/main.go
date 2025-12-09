package main

import "fmt"

func main() {
	input := "dvdf"
	result := lengthOfLongestSubstring(input)
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {

	var seenLetters []string

	currLetter := ""
	alrHas := false
	currLongest := 0
	if len(s) == 1 {
		return 1
	}
	for i := 0; i != len(s); i++ {
		alrHas = false
		currLetter = string(s[i])
		if len(seenLetters) == 0 {
			seenLetters = append(seenLetters, currLetter)
			continue
		} else {
			for j := 0; j != len(seenLetters); j++ {
				if currLetter == seenLetters[j] {
					alrHas = true

					break
				}
			}
		}

		if alrHas {
			currLongest = max(currLongest, len(seenLetters))
			fmt.Println(seenLetters)
			seenLetters = nil
			seenLetters = append(seenLetters, currLetter)

		} else {
			seenLetters = append(seenLetters, currLetter)
			fmt.Println(seenLetters)
			if i == len(s)-1 {
				currLongest = max(currLongest, len(seenLetters))
				seenLetters = nil
				return currLongest
			}
		}
		if i == len(s)-1 {
			currLongest = max(currLongest, len(seenLetters))
			seenLetters = nil
			return currLongest
		}

	}
	return currLongest
}
