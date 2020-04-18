package __to14
// Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
//
// Example:
// Given a binary tree
//
// 		1
//	   / \
//	  2   3
//	/ \
// 4   5
//
// Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
//
// Note: The length of path between two nodes is represented by the number of edges between them.

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	dfs(root, &ans)
	return ans

}


func dfs(root *TreeNode, ans *int) int {
	if root == nil { return 0 }
	L := dfs(root.Left, ans)
	R := dfs(root.Right, ans)
	*ans = max(*ans, L + R)
	return max(L, R) + 1
}

func max(a, b int) int {
	if a > b { return a }
	return b
}