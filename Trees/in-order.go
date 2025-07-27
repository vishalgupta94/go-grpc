/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	left := inorderTraversal(root.Left)
	left = append(left, root.Val)
	right := inorderTraversal(root.Right)
	left = append(left, right...)
	return left

}

func inorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{}
	output := []int{}
	for root != nil || len(stack) > 0 {

		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			output = append(output, top.Val)
			root = top.Right
		}

	}
	return output

}
