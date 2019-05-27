package leetcode

import (
	"math"
)

// https://leetcode-cn.com/problems/divide-two-integers/
/**
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
说明:

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
*/

// 主要思路就是对于除数每次x2，通过位移实现，同步统计除数的数量，
// 在这个值再次位移时候大于被除数时候结束，然后被除数减去这个值继续重复，
// 直到不能再接近位置
func divide(dividend int, divisor int) int {
	m, n, res := dividend, divisor, 0

	if m < 0 {
		m = -m
	}
	if n < 0 {
		n = -n
	}

	for m >= n {
		t, p := n, 1
		for m > (t << 1) {
			t <<= 1
			p <<= 1
		}
		res += p
		m -= t
	}

	if (dividend < 0) != (divisor < 0) {
		res = -res
	}

	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}
