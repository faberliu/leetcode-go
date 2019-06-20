/*
Problem 1. Tow Sum

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

/*
#### 解题思路
- 遍历数组，用一个map储存已经出现过的元素
- 计算目标值和元素值的差是否在已出现在之前的元素中，若已出现，返回两个值对应的索引
- 若一直到遍历结束都未匹配成功，则数组中不存在加起来等于目标值的两个数
- 时间复杂度 O(n), 空间复杂度O(n)
*/

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
