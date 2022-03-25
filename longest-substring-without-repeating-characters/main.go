package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	runeMap := make(map[rune]int, 0)
	currentLength := 0
	maxLength := 0
	startIndex := 0
	for idx, c := range s {
		if runeIndex, ok := runeMap[c]; ok && runeIndex >= startIndex {
			startIndex = runeMap[c]
			runeMap[c] = idx
			currentLength = idx - startIndex
		} else {
			runeMap[c] = idx
			currentLength += 1
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}

func main() {
	input := "abcabcbb"
	res := lengthOfLongestSubstring(input)
	fmt.Printf("%v\n", res)
}
