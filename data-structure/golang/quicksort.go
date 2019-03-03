package dsa

// go写的一个粗制滥造的快排，可能不是很快吧...
func quicksort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	left, right := 0, n-1
	pivot := n / 2
	nums[right], nums[pivot] = nums[pivot], nums[right]
	for i, _ := range nums {
		if nums[i] < nums[right] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
	nums[right], nums[left] = nums[left], nums[right]
	quicksort(nums[:left])
	quicksort(nums[left+1:])
	return nums
}
