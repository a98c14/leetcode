package main

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	input := "abcabcbb"
	res := lengthOfLongestSubstring(input)
	fmt.Printf("%v\n", res)
	if res != 3 {
		t.Fail()
	}

	input2 := "bbbbb"
	res2 := lengthOfLongestSubstring(input2)
	fmt.Printf("%v\n", res2)
	if res2 != 1 {
		t.Fail()
	}

	input3 := "aab"
	res3 := lengthOfLongestSubstring(input3)
	fmt.Printf("%v\n", res3)
	if res3 != 2 {
		t.Fail()
	}

	input4 := "dvdf"
	res4 := lengthOfLongestSubstring(input4)
	fmt.Printf("%v\n", res4)
	if res4 != 3 {
		t.Fail()
	}

	input5 := "pwwkew"
	res5 := lengthOfLongestSubstring(input5)
	fmt.Printf("%v\n", res5)
	if res5 != 3 {
		t.Fail()
	}
}
