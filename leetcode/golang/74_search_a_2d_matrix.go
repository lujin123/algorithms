package leetcode

// https://leetcode-cn.com/problems/search-a-2d-matrix/
/**
74. 搜索二维矩阵

编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
示例 1:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
示例 2:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false
*/

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])
	if n == 0 {
		return false
	}

	row := -1

	for i := 0; i < m; i++ {
		if target == matrix[i][n-1] {
			return true
		}
		if target < matrix[i][n-1] {
			row = i
			break
		}
	}

	if row == -1 {
		return false
	}

	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		if target == matrix[row][mid] {
			return true
		} else if target > matrix[row][mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
