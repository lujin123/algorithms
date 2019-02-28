package leetcode

import (
	"bytes"
	"strings"
)

// ref: https://leetcode-cn.com/problems/simplify-path/

// 以 Unix 风格给出一个文件的绝对路径，你需要简化它。或者换句话说，将其转换为规范路径。

// 在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..） 表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成部分。更多信息请参阅：Linux / Unix中的绝对路径 vs 相对路径

// 请注意，返回的规范路径必须始终以斜杠 / 开头，并且两个目录名之间必须只有一个斜杠 /。最后一个目录名（如果存在）不能以 / 结尾。此外，规范路径必须是表示绝对路径的最短字符串。

// 示例 1：

// 输入："/home/"
// 输出："/home"
// 解释：注意，最后一个目录名后面没有斜杠。
// 示例 2：

// 输入："/../"
// 输出："/"
// 解释：从根目录向上一级是不可行的，因为根是你可以到达的最高级。
// 示例 3：

// 输入："/home//foo/"
// 输出："/home/foo"
// 解释：在规范路径中，多个连续斜杠需要用一个斜杠替换。
// 示例 4：

// 输入："/a/./b/../../c/"
// 输出："/c"
// 示例 5：

// 输入："/a/../../b/../c//.//"
// 输出："/c"
// 示例 6：

// 输入："/a//b////c/d//././/.."
// 输出："/a/b/c"

// 思路：这个题目最开始没太懂他的”.“代表当前目录什么意思，后来发现就是当前目录可以不用处理，直接忽略...
// 1. 那就直接分割开用一个stack来保存，遇到..就出栈，栈空了就忽略，最后就拼接下就可以了，注意顺序
// 2. 直接调用系统的方法，例如golang的`filepath.Clean`，这个解法简直了

func simplifyPath(path string) string {
	s := strings.Split(path, "/")
	var stack []string
	for _, v := range s {
		if v == "" || v == "." {
			continue
		} else if v == ".." {
			n := len(stack)
			if n > 0 {
				stack = stack[:n-1]
			}
		} else {
			stack = append(stack, v)
		}
	}
	if len(stack) > 0 {
		var buffer bytes.Buffer
		for _, v := range stack {
			buffer.WriteString("/")
			buffer.WriteString(v)
		}
		return buffer.String()
	} else {
		return "/"
	}
}
