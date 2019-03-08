package leetcode

// https://leetcode-cn.com/problems/triangle/
/**
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

说明：

如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
*/

//自底向上动态规划, dp[i]表示到第i行时的最小路径和, 先从底层开始判断, 每向上一层后该层的数据无需再使用可以直接覆盖, 所以可以把空间复杂度优化到O(N), 但是要注意覆盖时的顺序
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n < 1 {
		return 0
	}
	dp := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		t := triangle[i]
		for j := 0; j < len(t); j++ {
			if dp[j] < dp[j+1] {
				dp[j] = dp[j] + t[j]
			} else {
				dp[j] = dp[j+1] + t[j]
			}
		}
	}

	return dp[0]
}
