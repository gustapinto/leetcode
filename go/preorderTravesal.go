package main

// https://leetcode.com/problems/binary-tree-preorder-traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var preorder func(root *TreeNode, carry []int)
	preorder = func(root *TreeNode, carry []int) {
		if root != nil {
			result = append(carry, root.Val)

			preorder(root.Left, result)
			preorder(root.Right, result)
		}
	}

	preorder(root, []int{})

	return result
}
