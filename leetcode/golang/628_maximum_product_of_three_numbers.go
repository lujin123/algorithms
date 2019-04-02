package leetcode

import (
	"math"
)

// https://leetcode-cn.com/problems/maximum-product-of-three-numbers/
/**
给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

示例 1:

输入: [1,2,3]
输出: 6
示例 2:

输入: [1,2,3,4]
输出: 24

注意:

1. 给定的整型数组长度范围是[3,104]，数组中所有的元素范围是[-1000, 1000]。
2. 输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。
*/

// 三个数乘积的最大值只会出现的最小的两个数和最大数的乘积或者最大三个数的乘积，可以找出最大的三个数和最小的两个数，判断下取最大值即可
func maximumProduct(nums []int) int {
	max1 := -1000
	max2 := -1000
	max3 := -1000
	min1 := 1000
	min2 := 1000

	ans := math.MinInt32

	for _, num := range nums {
		if num > max1 {
			max1, max2, max3 = num, max1, max2
		} else if num > max2 {
			max2, max3 = num, max2
		} else if num > max3 {
			max3 = num
		}

		if num < min1 {
			min1, min2 = num, min1
		} else if num < min2 {
			min2 = num
		}
	}
	a1 := max1 * max2 * max3
	a2 := min1 * min2 * max1
	if a1 > a2 {
		ans = a1
	} else {
		ans = a2
	}

	return ans
}
