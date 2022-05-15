package easy

import (
	"testing"
)

func TestMaxDepth1(t *testing.T) {
	var tree1 = TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	var result = MaxDepth(&tree1)
	t.Logf("[3,9,20,nil,nil,15,7] %v", result)
	if result != 3 {
		t.Fail()
	}
}

func TestMaxDepth2(t *testing.T) {
	var tree1 = TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	var result = MaxDepth(&tree1)
	t.Logf("[1,2,3,4,nil,nil,5] %v", result)
	if result != 3 {
		t.Fail()
	}
}
