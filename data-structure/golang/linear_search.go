package dsa

// 线性查找
func linearSearch(arr []int, value int) bool {
	for i := 0; i < len(arr); i++ {
		if value == arr[i] {
			return true
		}
	}
	return false
}

// 线性查找改进，返回查找的index,从两次比较减少为一次比较
func linearSearch2(arr []int, value int) int {
	var index int

	if arr[0] == value {
		return 0
	}
	arr[0] = value
	for index = len(arr) - 1; arr[index] != arr[0]; index-- {
	}
	if index == 0 {
		return -1
	}
	return index
}
