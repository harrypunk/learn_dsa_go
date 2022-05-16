package easy

import (
	"math"
)

func IsValidBST(root *TreeNode) bool {
	current := root
	stack := []*TreeNode{}
	pre := math.MinInt64
	for len(stack) > 0 || current != nil {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.Val <= pre {
			return false
		}
		pre = current.Val
		current = current.Right
	}
	return true
}
