/*
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:

Input: [3,2,3]
Output: 3
Example 2:

Input: [2,2,1,1,1,2,2]
Output: 2
*/
package main

import (
	"sort"
)

/*
#### 解题思路1
- 排序再遍历统计数量，当谁的数量超过n/2时，返回即可
*/
func majorityElement1(nums []int) int {
	sort.Ints(nums)
	count := 0
	for i := range nums {
		if i > 0 && nums[i] != nums[i-1] {
			count = 1
		} else {
			count++
		}
		if count > len(nums)/2 {
			return nums[i]
		}
	}
	return 0
}

/*
#### 解题思路2
- 排序，返回第n/2个元素即可(因为众数必定存在且大于n/2)
*/
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

/*
#### 解题思路3
- 分治
- 递归获取左右两边出现次数最多的数，如果相等，就是该数，否则，返回两边出现次数高的数
*/
func majorityElement3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := majorityElement3(nums[0 : len(nums)/2])
	right := majorityElement3(nums[len(nums)/2:])

	if left == right {
		return left
	}

	leftCount := countMajority(nums, left)
	rightCount := countMajority(nums, right)

	if leftCount > rightCount {
		return left
	}
	return right
}

func countMajority(nums []int, num int) int {
	count := 0
	for _, v := range nums {
		if v == num {
			count++
		}
	}
	return count
}

/*
#### 解题思路4
- Boyer-Moore 投票算法
- 遍历数组，当count为0时，该遍历到的数作为候选众数，继续往后，如果和候选众数一致，count+1，反之，count-1
- 当count=0时，重置候选众数。本质上是众数和非众数1，1消除，由于众数>n/2，所以最后剩下的一定是众数
*/
func majorityElement4(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	majority, count := nums[0], 0
	for i := range nums {
		if nums[i] != majority {
			count--
			if count == 0 {
				majority = nums[i]
				count = 1
			}
		} else {
			count++
		}
	}
	return majority
}
