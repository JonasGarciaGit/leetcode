package main

import "fmt"

//Given the root of a binary tree, invert the tree, and return its root.
// Input: root = [4,2,7,1,3,6,9]
// Output: [4,7,2,9,6,3,1]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	invertTree(root.Left)
	invertTree(root.Right)

	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	return root
}

func main() {
	treeNode := TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}

	fmt.Println(invertTree(&treeNode))
}
