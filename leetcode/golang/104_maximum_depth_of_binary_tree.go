package leetcode

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

/**
104. 二叉树的最大深度

给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 直接递归实现
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	n1 := maxDepth(root.Left)
	n2 := maxDepth(root.Right)
	if n1 > n2 {
		return n1 + 1
	}
	return n2 + 1
}

// 利用队列存储每一层的节点，每遍历一层就+1并且将子节点加入队列，直到队列为空
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	queue = append(queue, root)
	ans := 0
	n := len(queue)
	for n > 0 {
		ans++
		var temp []*TreeNode
		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				temp = append(temp, queue[i].Left)
			}
			if queue[i].Right != nil {
				temp = append(temp, queue[i].Right)
			}
		}
		queue = temp
		n = len(queue)
	}
	return ans
}
