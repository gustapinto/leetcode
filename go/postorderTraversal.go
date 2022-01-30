package main

// https://leetcode.com/problems/binary-tree-preorder-traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)

    var postorder func(node *TreeNode)
    postorder = func(node *TreeNode) {
        if node != nil {
            postorder(node.Left)
            postorder(node.Right)

            result = append(result, node.Val)
        }
    }

    postorder(root)

    return result
}
