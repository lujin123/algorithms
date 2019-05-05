package leetcode

import (
	"strconv"
)

// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
/**
150. 逆波兰表达式求值

根据逆波兰表示法，求表达式的值。

有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

说明：

整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
示例 1：

输入: ["2", "1", "+", "3", "*"]
输出: 9
解释: ((2 + 1) * 3) = 9
示例 2：

输入: ["4", "13", "5", "/", "+"]
输出: 6
解释: (4 + (13 / 5)) = 6
示例 3：

输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
输出: 22
解释:
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
*/

// 逆波兰表达式求值过程就是利用栈存储数字，然后遇到操作符，二元操作符则弹出两个字符之后求值即可
func evalRPN(tokens []string) int {
	var stack []int
	for _, v := range tokens {
		n := len(stack)
		var res int
		switch v {
		case "+":
			res = stack[n-2] + stack[n-1]
		case "-":
			res = stack[n-2] - stack[n-1]
		case "*":
			res = stack[n-2] * stack[n-1]
		case "/":
			res = stack[n-2] / stack[n-1]
		default:
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
			continue
		}
		stack = stack[:n-2]
		stack = append(stack, res)
	}
	return stack[0]
}
