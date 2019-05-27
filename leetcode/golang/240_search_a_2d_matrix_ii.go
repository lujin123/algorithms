package leetcode

// https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
/**
240. 搜索二维矩阵 II

编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。
*/

/*
左下角的元素是这一行中最小的元素，同时又是这一列中最大的元素。比较左下角元素和目标：
若左下角元素等于目标，则找到
若左下角元素大于目标，则目标不可能存在于当前矩阵的最后一行，问题规模可以减小为在去掉最后一行的子矩阵中寻找目标
若左下角元素小于目标，则目标不可能存在于当前矩阵的第一列，问题规模可以减小为在去掉第一列的子矩阵中寻找目标
若最后矩阵减小为空，则说明不存在
*/
func searchMatrixII(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	i := m - 1
	j := 0
	for i >= 0 && j < n {
		if target == matrix[i][j] {
			return true
		}
		if target > matrix[i][j] {
			j++
		} else {
			i--
		}
	}
	return false
}
