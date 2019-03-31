package leetcode

// https://leetcode-cn.com/problems/add-two-numbers/

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 简单粗暴的创建一个新的链表，然后注意长度和进位即可
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	curry := 0
	var header *ListNode
	var p *ListNode
	for p1 != nil && p2 != nil {
		val := p1.Val + p2.Val + curry
		if val > 9 {
			curry = 1
			val -= 10
		} else {
			curry = 0
		}
		node := &ListNode{
			Val: val,
		}
		if header == nil {
			header = node
			p = header
		} else {
			p.Next = node
			p = p.Next
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	var p3 *ListNode
	if p1 != nil {
		p3 = p1
	} else {
		p3 = p2
	}
	for p3 != nil {
		val := p3.Val + curry
		if val > 9 {
			curry = 1
			val -= 10
		} else {
			curry = 0
		}
		node := &ListNode{
			Val: val,
		}
		if header == nil {
			header = node
			p = header
		} else {
			p.Next = node
			p = p.Next
		}
		p3 = p3.Next
	}
	if curry != 0 {
		node := &ListNode{
			Val: curry,
		}
		p.Next = node
	}
	return header
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	header := l1
	var p *ListNode
	curry := 0
	for l1 != nil || l2 != nil {
		var (
			v1 int
			v2 int
		)
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}
		val := v1 + v2 + curry
		curry = val / 10
		val = val % 10

		var node *ListNode
		if l1 != nil {
			node = l1
		} else {
			node = l2
		}
		node.Val = val

		if p != nil {
			p.Next = node
		}
		p = node

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if curry != 0 {
		node := &ListNode{Val: curry}
		if p != nil {
			p.Next = node
		}
	}

	return header
}

func execAddTwoNumbers(nums1 []int, nums2 []int) []int {
	l1 := createList(nums1)
	l2 := createList(nums2)
	header := addTwoNumbers(l1, l2)
	res := printList(header)
	return res
}

func execAddTwoNumbers2(nums1 []int, nums2 []int) []int {
	l1 := createList(nums1)
	l2 := createList(nums2)
	header := addTwoNumbers2(l1, l2)
	res := printList(header)
	return res
}
