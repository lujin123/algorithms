package leetcode

// https://leetcode-cn.com/problems/sort-colors/
/**
75. 颜色分类

给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

示例:

输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]
进阶：

一个直观的解决方案是使用计数排序的两趟扫描算法。
首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
你能想出一个仅使用常数空间的一趟扫描算法吗？
*/

func sortColors(nums []int) {
	count0 := 0
	count1 := 0
	count2 := 0
	for i := range nums {
		switch nums[i] {
		case 0:
			count0++
		case 1:
			count1++
		case 2:
			count2++
		}
	}
	for i := 0; i < count0; i++ {
		nums[i] = 0
	}
	for i := count0; i < count0+count1; i++ {
		nums[i] = 1
	}
	for i := count0 + count1; i < len(nums); i++ {
		nums[i] = 2
	}
}

// 三路快排，将0全部换到左边，2换到右边，注意如果换的是右边的值，那么需要再次判断调换的值是否要再换
// 一遍扫描即可排完
func sortColors2(nums []int) {
	l := 0
	r := len(nums)

	for i := 0; i < r; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
			i++
		} else {
			r--
			nums[r], nums[i] = nums[i], nums[r]
		}
	}
}
