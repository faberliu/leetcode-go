package main

/*
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:

[
  [3],
  [9,20],
  [15,7]
]
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
#### 解题思路1
- 使用广度优先遍历，将每层遍历结果输出即可
*/

func levelOrder1(root *TreeNode) [][]int {
	var result [][]int
	var queue []*TreeNode
	if root == nil {
		return nil
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		var level []int
		var nodes []*TreeNode
		for _, v := range queue {
			level = append(level, v.Val)
			if v.Left != nil {
				nodes = append(nodes, v.Left)
			}
			if v.Right != nil {
				nodes = append(nodes, v.Right)
			}

		}
		queue = nodes
		result = append(result, level)
	}
	return result
}

/*
#### 解题思路2
- 深度优先遍历
*/
func levelOrder2(root *TreeNode) [][]int {
	var result [][]int
	DFSOrder(root, 0, &result)
	return result
}

func DFSOrder(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}
	if len(*result) < level+1 {
		*result = append(*result, []int{root.Val})
	} else {
		(*result)[level] = append((*result)[level], root.Val)
	}
	DFSOrder(root.Left, level+1, result)
	DFSOrder(root.Right, level+1, result)
}
