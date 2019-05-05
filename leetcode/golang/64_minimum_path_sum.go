package leetcode

// https://leetcode-cn.com/problems/minimum-path-sum/
/**
64. 最小路径和

给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/

// 这题用递归暴力求解会超时，所以需要用到动态规划

// 动态规划，直接新建一个数组保存每次计算的值，上面和左面的两排数据直接累加即可，
// 中间部分需要计算上面和左面的最小值
// 时间复杂度 = O(m*n)，空间复杂度 = O(m*n)
func minPathSum(grid [][]int) int {
	if grid == nil {
		return 0
	}
	row := len(grid)
	col := len(grid[0])

	var dp [][]int
	for i := 0; i < row; i++ {
		dp = append(dp, make([]int, col))
	}

	var temp1 int
	for i := 0; i < row; i++ {
		dp[i][0] = temp1 + grid[i][0]
		temp1 = dp[i][0]
	}
	var temp2 int
	for i := 0; i < col; i++ {
		dp[0][i] = temp2 + grid[0][i]
		temp2 = dp[0][i]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			min := dp[i-1][j]
			if dp[i][j-1] < min {
				min = dp[i][j-1]
			}
			dp[i][j] = min + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

// 优化下空间复杂度，直接用原来的数组保存，不用新开空间
// 时间复杂度 = O(m*n)，空间复杂度 = O(1)
func minPathSum2(grid [][]int) int {
	if grid == nil {
		return 0
	}
	row := len(grid)
	col := len(grid[0])

	var temp1 int
	for i := 0; i < row; i++ {
		grid[i][0] += temp1
		temp1 = grid[i][0]
	}
	var temp2 int
	for i := 0; i < col; i++ {
		grid[0][i] += temp2
		temp2 = grid[0][i]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			min := grid[i-1][j]
			if grid[i][j-1] < min {
				min = grid[i][j-1]
			}
			grid[i][j] += min
		}
	}
	return grid[row-1][col-1]
}
