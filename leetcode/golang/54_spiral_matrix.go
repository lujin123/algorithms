package leetcode

// https://leetcode-cn.com/problems/spiral-matrix/
/**
54. 螺旋矩阵

给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

func spiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}
	m := len(matrix)
	if m < 1 {
		return []int{}
	}
	n := len(matrix[0])

	var (
		row int
		col int
		res []int
	)
	for row < m && col < n {
		for i := col; i < n; i++ {
			res = append(res, matrix[row][i])
		}

		for i := row + 1; i < m; i++ {
			res = append(res, matrix[i][n-1])
		}

		for i := n - 2; i >= col && m-1 > row; i-- {
			res = append(res, matrix[m-1][i])
		}

		for i := m - 2; i > row && col < n-1; i-- {
			res = append(res, matrix[i][col])
		}

		col++
		row++
		n--
		m--
	}
	return res
}
