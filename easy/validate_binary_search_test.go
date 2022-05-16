package easy

import (
	"testing"
)

func TestIsValidBST1(t *testing.T) {
	root := TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	var result = IsValidBST(&root)
	if !result {
		t.Fail()
	}
}
