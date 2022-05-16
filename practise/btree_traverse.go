package practise

import (
	"dsa_go/easy"
	"fmt"
)

func Preorder(root *easy.TreeNode) {
	stack := []*easy.TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			fmt.Printf("%v,", current.Val)

			stack = append(stack, current)
			current = current.Left
		}

		var last = len(stack)
		if last > 0 {
			current = stack[last-1]
			stack = stack[:last-1]
			current = current.Right
		}
	}
  fmt.Println()
}
