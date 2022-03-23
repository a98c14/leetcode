package main

func twoSumV2(nums []int, target int) []int {
	numberMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		if k, ok := numberMap[target-a]; ok {
			return []int{k, i}
		}
		numberMap[a] = i
	}

	return nil
}
