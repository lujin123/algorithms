package dsa

// 二分查找,前提是排好序的数组
func binarySearch(arr []int, value int) bool {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == value {
			return true
		} else if value > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
