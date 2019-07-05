/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

/*
#### 解题思路1
- 遍历nums
- 对于每一个元素a，遍历其它元素b，看是否存在一个c使 a+b+c = 0
- 如果存在，加入到res中，如果不存在，继续下一个
- 为了避免返回重复结果，再增加一个去重判断
*/
package main

import (
	"sort"
)

func threeSum1(nums []int) [][]int {
	var res [][]int
	resSet := make(map[[3]int]struct{})
	if len(nums) <= 0 {
		return res
	}
	sort.Ints(nums)
	last := nums[0]
	if last > 0 {
		return res
	}
	last = 1
	for i, a := range nums {
		if last == a {
			continue
		}
		last = a
		dict := make(map[int]int)
		for j, b := range nums[i+1:] {
			if _, ok := dict[-(a + b)]; ok {
				a := [3]int{a, -(a + b), b}
				if _, ok := resSet[a]; !ok {
					res = append(res, a[:])
					resSet[a] = struct{}{}
				}
			} else {
				dict[b] = j
			}
		}
	}
	return res
}

/*
#### 解题思路2
- 排序
- 第一层采用循环，对于每个a，对剩余的数字采用左右逼近
- 若左右之和大于-a，则右边退一格，若左右之和小于-a，则左边前进一格
- 若某边重复,则直接跳过
*/

func threeSum2(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			total := nums[left] + nums[right] + nums[i]
			if total > 0 {
				right--
			} else if total < 0 {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				right--
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return res
}
