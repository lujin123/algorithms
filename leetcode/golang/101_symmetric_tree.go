package leetcode

// https://leetcode-cn.com/problems/symmetric-tree/

/**
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
说明:

如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
*/

// 1. 递归解法
func isSymmetric(root *TreeNode) bool {
	return judge(root, root)
}

func judge(p1 *TreeNode, p2 *TreeNode) bool {
	if p1 == nil && p2 == nil {
		return true
	}
	if p1 == nil || p2 == nil {
		return false
	}
	return p1.Val == p2.Val && judge(p1.Left, p2.Right) && judge(p1.Right, p2.Left)
}

// 2. 迭代法，用队列保存节点，每两个连续节点应该相等
func isSymmetric2(root *TreeNode) bool {
	var queue []*TreeNode
	// 先将两个根节点放进去
	queue = append(queue, root, root)
	index := 0
	for index < len(queue) {
		left := queue[index]
		right := queue[index+1]
		index += 2
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}
