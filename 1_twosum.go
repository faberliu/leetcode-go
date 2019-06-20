package main

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, v := range nums {
		if m, ok := dict[target-v]; ok {
			return []int{m, i}
		}
		dict[v] = i
	}
	return nil
}
