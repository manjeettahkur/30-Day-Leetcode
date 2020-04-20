package _5_to_21

//Return the root node of a binary search tree that matches the given preorder traversal.
//
//(Recall that a binary search tree is a binary tree where for every node, any descendant of node.left has a value < node.val, and any descendant of node.right has a value > node.val.  Also recall that a preorder traversal displays the value of the node first, then traverses node.left, then traverses node.right.)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	l := len(preorder)
	return InsertNode(preorder, 0, l)
}

func InsertNode(preorder []int, rootIndex int, l int) *TreeNode {

	if rootIndex >= l {
		return nil
	}

	value := preorder[rootIndex]

	root := &TreeNode{Val: value}
	i := rootIndex + 1
	for i <= l-1 && preorder[i] < value {
		i++
	}

	root.Left = InsertNode(preorder, rootIndex+1, i)
	root.Right = InsertNode(preorder, i, l)
	return root
}
