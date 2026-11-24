package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Identical reports whether two trees have the same structure and values.
func Identical(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b // both must be nil together
	}
	return a.Val == b.Val && Identical(a.Left, b.Left) && Identical(a.Right, b.Right)
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	a := tn(1, tn(2, nil, nil), tn(3, nil, nil))
	b := tn(1, tn(2, nil, nil), tn(3, nil, nil))
	c := tn(1, nil, tn(2, nil, nil))
	fmt.Println(Identical(a, b), Identical(a, c))
}
