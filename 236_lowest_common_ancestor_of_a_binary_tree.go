package main

/*
236. Lowest Common Ancestor of a Binary Tree
Medium

2026

138

Favorite

Share
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Given the following binary tree:  root = [3,5,1,6,2,0,8,null,null,7,4]




Example 1:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.
Example 2:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.


Note:

All of the nodes' values will be unique.
p and q are different and both values will exist in the binary tree.
*/
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

/*
#### 解题思路1
- 在左右子树中分别查找是否包含p或q，
- 如果左子树包含p/q, 右子树包含另一个，则当前节点为公共节点,返回
- 如果右子树中不包含p和q，则公共节点必定在左子树中
- 如果左子树中不包含p和q，则公共节点必定在右子树中
*/
func lowestCommonAncestor_1(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}
	left := lowestCommonAncestor_1(root.Left, p, q)
	right := lowestCommonAncestor_1(root.Right, p, q)
	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}

/*
#### 解题思路2
- 采用深度优先遍历的方式，记录是否出现了p，q和中间节点，如果出现超过2个，则当前节点就是最小父节点
- 递归函数返回的是是否找到了其中一个或一个以上节点
*/
func lowestCommonAncestor_2(root, p, q *TreeNode) *TreeNode {
	var result *TreeNode
	findPQ(root, p, q, &result)
	return result
}

func findPQ(root, p, q *TreeNode, result **TreeNode) bool {
	if root == nil {
		return false
	}
	left, right, mid := 0, 0, 0
	if findPQ(root.Left, p, q, result) {
		left = 1
	}
	if findPQ(root.Right, p, q, result) {
		right = 1
	}
	if p == root || q == root {
		mid = 1
	}
	found := mid + right + left
	if found > 1 {
		*result = root
	}
	return found > 0
}
