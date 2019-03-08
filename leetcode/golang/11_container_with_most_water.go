package leetcode

// https://leetcode-cn.com/problems/container-with-most-water/
/**
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

[图片直接去连接原题去看]

图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。


示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49
*/

func maxArea(height []int) int {
	n := len(height)
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			h := minInt(height[i], height[j])
			dp[i] = maxInt((i-j)*h, dp[i])
		}
	}
	max := 0
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}

func maxArea2(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		area := 0
		if height[left] < height[right] {
			area = (right - left) * height[left]
			left++
		} else {
			area = (right - left) * height[right]
			right--
		}
		if area > max {
			max = area
		}
	}
	return max
}
