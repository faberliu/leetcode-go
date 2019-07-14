package main

import "math"

/*
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth1(root.Left)
	rightMax := maxDepth1(root.Right)
	if leftMax > rightMax {
		return leftMax + 1
	}
	return rightMax + 1
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + (int)(math.Max((float64)(maxDepth2(root.Left)), float64(maxDepth2(root.Right))))
}
