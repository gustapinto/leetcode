package main

// https://leetcode.com/problems/binary-tree-inorder-traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)

			result = append(result, node.Val)

			inorder(node.Right)
		}
	}

	inorder(root)

	return result
}
