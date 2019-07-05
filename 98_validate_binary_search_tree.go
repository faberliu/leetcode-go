package main

import (
	"math"
	"reflect"
	"sort"
)

/*
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.


Example 1:

    2
   / \
  1   3

Input: [2,1,3]
Output: true
Example 2:

    5
   / \
  1   4
     / \
    3   6

Input: [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
#### 解题思路1
- 中序遍历，比较中序遍历后的值是否有序
*/

func isValidBST_1(root *TreeNode) bool {
	inOrderSlice := inOrder(root)
	inOrderSliceOrdered := distict(inOrderSlice)
	sort.Ints(inOrderSliceOrdered)
	return reflect.DeepEqual(inOrderSlice, inOrderSliceOrdered)
}

func distict(s []int) []int {
	var res []int
	_map := make(map[int]struct{})
	for _, v := range s {
		if _, ok := _map[v]; !ok {
			res = append(res, v)
			_map[v] = struct{}{}
		}
	}
	return res
}

func inOrder(root *TreeNode) []int {
	var res []int
	if root == nil {
		return nil
	}
	res = append(res, inOrder(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inOrder(root.Right)...)
	return res
}

/*
#### 解题思路2
- 中序遍历，依次比较和前一个元素的大小，如果都大于前一个元素，则合法
*/
func isValidBST_2(root *TreeNode) bool {
	var pre *TreeNode = nil
	return isBSTHelper(&pre, root)
}

func isBSTHelper(pre **TreeNode, root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBSTHelper(pre, root.Left) {
		return false
	}
	if *pre != nil && (*pre).Val >= root.Val {
		return false
	}
	*pre = root
	return isBSTHelper(pre, root.Right)
}

/*
#### 解题思路3
- 递归，左子树最大值必须<根<右子树最小值
*/
func isValidBST_3(root *TreeNode) bool {
	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}

	if left >= root.Val || right <= root.Val {
		return false
	}

	return isBST(root.Left, left, root.Val) && isBST(root.Right, root.Val, right)
}
