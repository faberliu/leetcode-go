package main

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

/*
#### 解题思路
- 递归构建，基于已经构建好的合法序列去构建新的序列
- left表示已经有多少个左括号
- right表示已经有多少个右括号
- 每次构建有两种选择，左括号或者右括号
- 当left小于N时，可以选择左括号
- 当right小于left时，可以选择右括号(否则不是合法序列)
- 当构建出长度为2N的括号时，即为其中一个答案，加入结果中
*/
func generateParenthesis(n int) []string {
	var result []string
	var generate func(s string, left, right int)
	generate = func(s string, left, right int) {
		if len(s) == 2*n {
			result = append(result, s)
		}
		if left < n {
			generate(s+"(", left+1, right)
		}
		if right < left {
			generate(s+")", left, right+1)
		}
	}
	generate("", 0, 0)
	return result
}
