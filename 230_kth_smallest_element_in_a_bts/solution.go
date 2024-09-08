package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	slice := []int{}
	kthSmallestHelper(root, &slice)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice[k-1]
}

func kthSmallestHelper(root *TreeNode, slice *[]int) {
	if root == nil {
		return
	}

	*slice = append(*slice, root.Val)
	kthSmallestHelper(root.Left, slice)
	kthSmallestHelper(root.Right, slice)
}

func main() {
	treeNode := TreeNode{
		3,
		&TreeNode{
			1,
			nil,
			&TreeNode{
				2,
				nil,
				nil,
			},
		},
		&TreeNode{
			4,
			nil,
			nil,
		},
	}
	fmt.Println(kthSmallest(&treeNode, 1))
}
