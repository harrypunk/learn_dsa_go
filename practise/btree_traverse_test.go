package practise

import (
	. "dsa_go/easy"
	"testing"
)

func TestPreorder(t *testing.T) {
	var root = TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	var result = Preorder(&root)
	t.Logf("%v", result)
}
