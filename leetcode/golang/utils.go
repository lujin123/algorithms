package leetcode

func minInt(v1, v2 int) int {
	if v1 > v2 {
		return v2
	}
	return v1
}

func maxInt(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
