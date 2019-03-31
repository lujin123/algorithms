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
