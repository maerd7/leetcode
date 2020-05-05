package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
	  *     Val int
	   *     Left *TreeNode
	    *     Right *TreeNode
		 *
 }
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getBothDepth(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	if node.Left == nil && node.Right == nil {
		return 1, true
	}
	left, leftStatus := getBothDepth(node.Left)
	right, rightStatus := getBothDepth(node.Right)
	if !leftStatus || !rightStatus {
		return 0, false
	}
	if left > right {
		if left-right > 1 {
			return left + 1, false
		} else {
			return left + 1, true
		}
	} else {
		if right-left > 1 {
			return right + 1, false
		} else {
			return right + 1, true
		}
	}
}

// 5.3 15:00 ending time 5.3 16:53
func isBalanced(root *TreeNode) bool {
	_, status := getBothDepth(root)
	return status
}
