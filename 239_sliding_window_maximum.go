/*
Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.

Example:

Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7]
Explanation:

Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Note:
You may assume k is always valid, 1 ≤ k ≤ input array's size for non-empty array.

Follow up:
Could you solve it in linear time?
*/
/*
解体思路
- 用一个window队列储存窗口对应的index
- 每滑动一个窗口，进行如下操作
	- 判断窗口中第一个元素是否已超出窗口范围，如果是删除第一个元素
	- 从右往左判断窗口中元素是否小于新增元素，如果是，删除
	- 将新元素增加到窗口中
	- 将第一个元素(最大元素)追加到结果队列
*/
package main

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) <= 0 {
		return []int{}
	}
	var (
		window []int
		res    []int
	)
	for i, v := range nums {
		if i > k-1 && window[0] < i+1-k {
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < v {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}

/*
func main() {
	test := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(test, k))
}*/
