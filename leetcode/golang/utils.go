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

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printList(head *ListNode) []int {
	var res []int
	p := head
	for p != nil {
		res = append(res, p.Val)
		p = p.Next
	}
	return res
}

func createList(nums []int) *ListNode {
	var head *ListNode
	var p *ListNode
	for _, val := range nums {
		node := &ListNode{
			Val: val,
		}
		if head == nil {
			head = node
			p = head
		} else {
			p.Next = node
			p = node
		}
	}
	return head
}

func createFulllikeTreeWithArray(nums []int) (root *TreeNode) {
	n := len(nums)
	if n == 0 {
		return
	}
	var queue []*TreeNode

	index := 0

	queue = append(queue, &TreeNode{
		Val: nums[index],
	})

	index += 1

	for len(queue) > 0 && index < n {
		node := queue[0]
		queue = queue[1:]
		if root == nil {
			root = node
		}
		leftVal := nums[index]
		if leftVal != 0 {
			leftNode := &TreeNode{Val: leftVal}
			node.Left = leftNode
			queue = append(queue, leftNode)
		}
		index += 1

		if index >= n {
			break
		}
		rightVal := nums[index]
		if rightVal != 0 {
			rightNode := &TreeNode{Val: rightVal}
			node.Right = rightNode
			queue = append(queue, rightNode)
		}
		index += 1
	}
	return
}
