package practise

import (
	"dsa_go/easy"
	"testing"
)

func TestPreorder(t *testing.T) {
	var root = easy.TreeNode{
		Val: 12,
	}

	Preorder(&root)
}
