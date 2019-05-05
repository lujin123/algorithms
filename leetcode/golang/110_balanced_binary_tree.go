package leetcode

// https://leetcode-cn.com/problems/balanced-binary-tree/
/**
110. 平衡二叉树

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	n1 := getDepth(root.Left)
	n2 := getDepth(root.Right)

	delta := n1 - n2
	if delta > 1 || delta < -1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	n1 := getDepth(root.Left)
	n2 := getDepth(root.Right)
	if n1 > n2 {
		return n1 + 1
	}
	return n2 + 1
}
