package leetcode

import (
	"math"
	"strconv"
)

// https://leetcode-cn.com/problems/basic-calculator/
/**
224. 基本计算器

实现一个基本的计算器来计算一个简单的字符串表达式的值。

字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，非负整数和空格  。

示例 1:

输入: "1 + 1"
输出: 2
示例 2:

输入: " 2-1 + 2 "
输出: 3
示例 3:

输入: "(1+(4+5+2)-3)+(6+8)"
输出: 23
说明：

你可以假设所给定的表达式都是有效的。
请不要使用内置的库函数 eval。
*/
// 主要思路就是：
// 1. 先将中缀表达式转换成一个后缀表达式,这一步需要注意多个数字的计算
// 2. 利用辅助栈来计算后缀表达式的值
func calculate(s string) int {
	res := convert2postfix(s)
	var stack []int
	for i := range res {
		switch res[i] {
		case "+":
			n := len(stack)
			sum := stack[n-1] + stack[n-2]
			stack = stack[:n-2]
			stack = append(stack, sum)
		case "-":
			n := len(stack)
			sub := stack[n-2] - stack[n-1]
			stack = stack[:n-2]
			stack = append(stack, sub)
		default:
			stack = append(stack, str2int(res[i]))
		}
	}
	return stack[0]
}

func convert2postfix(s string) []string {
	var (
		res   []string
		stack []rune
	)
	var numbers []rune
	for _, v := range s {
		flag := isDigital(v)
		if flag {
			numbers = append(numbers, v)
		} else {
			if len(numbers) > 0 {
				res = append(res, string(numbers))
				numbers = numbers[:0]
			}
			switch v {
			case ' ':
				continue
			case '(':
				stack = append(stack, v)
			case ')':
				j := len(stack) - 1
				for ; j >= 0; j-- {
					if stack[j] == '(' {
						break
					}
					res = append(res, string(stack[j]))
				}
				stack = stack[:j]
			case '+', '-':
				nowPriority := getPriority(v)
				j := len(stack) - 1
				for ; j >= 0; j-- {
					priority := getPriority(stack[j])
					if nowPriority < priority {
						break
					}
					res = append(res, string(stack[j]))
				}
				stack = stack[:j+1]
				stack = append(stack, v)
			}
		}
	}

	if len(numbers) > 0 {
		res = append(res, string(numbers))
	}
	for i := len(stack) - 1; i >= 0; i-- {
		res = append(res, string(stack[i]))
	}
	return res
}

func getPriority(v rune) int {
	switch v {
	case '+', '-':
		return 1
	case '(', ')':
		return math.MaxInt32
	default:
		return 0
	}
}

func isDigital(v rune) bool {
	return v >= 48 && v <= 57
}

func str2int(v string) int {
	r, _ := strconv.Atoi(v)
	return r
}
