/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	var output [][]int
	var postTraversal func(*TreeNode, int)

	postTraversal = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(output) <= level {
			output = append(output, []int{})
		}

		output[level] = append(output[level], root.Val)

		postTraversal(root.Left, level+1)
		postTraversal(root.Right, level+1)
	}
	postTraversal(root, 0)

	return output

}
 
 