package leetcode

// https://leetcode-cn.com/problems/spiral-matrix-ii/
/**
59. 螺旋矩阵 II

给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
*/

func generateMatrix(n int) [][]int {
	var ans [][]int

	for i := 0; i < n; i++ {
		ans = append(ans, make([]int, n))
	}

	max := n * n
	var (
		row int
		col int
	)

	for i := 1; i <= max; {
		for j := col; j < n; j++ {
			ans[row][j] = i
			i++
		}
		for j := row + 1; j < n; j++ {
			ans[j][n-1] = i
			i++
		}

		for j := n - 2; j >= col && n-1 > row; j-- {
			ans[n-1][j] = i
			i++
		}

		for j := n - 2; j > row && n-1 > col; j-- {
			ans[j][col] = i
			i++
		}

		row++
		col++
		n--
	}
	return ans
}
