package practise

import (
	"dsa_go/easy"
)

func Preorder(root *easy.TreeNode) []int {
	res := []int{}
	stack := []*easy.TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			res = append(res, current.Val)

			stack = append(stack, current)
			current = current.Left
		}

		var last = len(stack)
		current = stack[last-1].Right
		stack = stack[:last-1]
	}
	return res
}

func PreorderRecur(root *easy.TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	res = append(res, root.Val)
	res = append(res, PreorderRecur(root.Left)...)
	res = append(res, PreorderRecur(root.Right)...)

	return res
}
