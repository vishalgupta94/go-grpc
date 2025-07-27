/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	output := []int{root.Val}
	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)
	output = append(output, left...)
	output = append(output, right...)
	return output
}

func preorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}
	output := []int{}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		output = append(output, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}

		if top.Left != nil {
			stack = append(stack, top.Left)
		}

	}
	return output
}
