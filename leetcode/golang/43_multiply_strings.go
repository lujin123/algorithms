package leetcode

import (
	"bytes"
	"strconv"
)

// doc: https://leetcode-cn.com/problems/multiply-strings/
// 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

// 示例 1:

// 输入: num1 = "2", num2 = "3"
// 输出: "6"
// 示例 2:

// 输入: num1 = "123", num2 = "456"
// 输出: "56088"
// 说明：

// num1 和 num2 的长度小于110。
// num1 和 num2 只包含数字 0-9。
// num1 和 num2 均不以零开头，除非是数字 0 本身。
// 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

// 思路：主要就是模拟乘法运算，然后保存每一位的值，之后再相加，这样可以避免溢出问题，也就是大数计算的方法，注意进位即可

func multiply(num1 string, num2 string) string {
	n1 := len(num1)
	n2 := len(num2)
	s1 := make([]int, n1)
	for index, v := range num1 {
		s1[index] = int(v - '0')
	}

	max := n1 + n2

	var results [][]int
	for i := n2 - 1; i >= 0; i-- {
		res := make([]int, max)
		v := int(num2[i] - '0')
		var curry int
		for j := n1 - 1; j >= 0; j-- {
			temp := v*s1[j] + curry
			if temp > 9 {
				curry = temp / 10
				temp = temp % 10
			} else {
				curry = 0
			}
			res[n1-j-1+n2-i-1] = temp
		}
		if curry > 0 {
			res[n1+n2-i-1] = curry
		}
		results = append(results, res)
	}
	// fmt.Printf("results: %v\n", results)

	var ss []int
	var curry int
	for i := 0; i < max; i++ {
		var temp int
		for j := 0; j < n2; j++ {
			temp += results[j][i]
		}
		temp += curry
		if temp > 9 {
			curry = temp / 10
			temp = temp % 10
		} else {
			curry = 0
		}
		ss = append(ss, temp)
	}
	if curry > 0 {
		for curry > 0 {
			curry = curry % 10
			ss = append(ss, curry)
		}
		ss = append(ss, curry)
	}
	// fmt.Println(ss)
	var buffer bytes.Buffer
	flag := false
	for i := len(ss) - 1; i >= 0; i-- {
		val := ss[i]
		if val == 0 && !flag {
			continue
		}
		buffer.WriteString(strconv.Itoa(val))
		flag = true
	}
	res := buffer.String()
	if res == "" {
		res = "0"
	}
	return res
}
