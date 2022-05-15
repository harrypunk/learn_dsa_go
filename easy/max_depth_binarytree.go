// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
package easy

func MaxDepth(root *TreeNode) int {
	//BFS
	if root == nil {
		return 0
	}

	max := 0
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		max++
		childs := []*TreeNode{}
		for i := 0; i < len(stack); i++ {
			var current = stack[i]
			if current.Right != nil {
				childs = append(childs, current.Right)
			}
			if current.Left != nil {
				childs = append(childs, current.Left)
			}
		}
		stack = childs
	}

	return max
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
