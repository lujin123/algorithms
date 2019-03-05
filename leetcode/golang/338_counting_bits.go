package leetcode

// https://leetcode-cn.com/problems/counting-bits/
/**
给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

示例 1:

输入: 2
输出: [0,1,1]
示例 2:

输入: 5
输出: [0,1,1,2,1,2]
进阶:

给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
要求算法的空间复杂度为O(n)。
你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）来执行此操作。
*/

/**
思路：
1. 保存每一个数的1的个数，先将数组与上1，得到最后一个是0或者1，然后将数字右移一位，这样去掉最后一位，得到新的数，从数组中找新的数的记录，之后加上前面与到的值即可
2. 或者将 i & (i-1)这样相当于将最后一位置为0，将这个结果取出的位数加一即可
*/

func countBits(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		dp[i] = (i & 1) + dp[i>>1]
	}
	return dp
}

func countBits2(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		dp[i] = dp[i&(i-1)] + 1
	}
	return dp
}
