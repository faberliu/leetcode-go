/*
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?
*/
package main

import (
	"reflect"
)

/*
#### 解决方案1
- 分别统计两个字符串中各字符的频次，比较频次是否相等
*/
func isAnagram1(s string, t string) bool {
	dict1, dict2 := make(map[rune]int), make(map[rune]int)
	for _, v := range s {
		dict1[v]++
	}
	for _, v := range t {
		dict2[v]++
	}
	return reflect.DeepEqual(dict1, dict2)
}

/*
#### 解决方案2
- 分别统计两个字符串中各字符的频次，保存在一个长度为26位的数组中
*/
func isAnagram2(s string, t string) bool {
	dict1, dict2 := make([]int, 26), make([]int, 26)
	for _, v := range s {
		dict1[v-'a']++
	}
	for _, v := range t {
		dict2[v-'a']++
	}
	return reflect.DeepEqual(dict1, dict2)
}
