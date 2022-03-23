package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		for k := i + 1; k < len(nums); k++ {
			b := nums[k]
			if a+b == target {
				return []int{i, k}
			}
		}
	}

	return nil
}

func main() {
	res := twoSumV2([]int{3, 2, 4}, 6)
	fmt.Printf(`%v`, res)
}
