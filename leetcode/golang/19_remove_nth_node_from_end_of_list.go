package leetcode

// doc: https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

// 示例：

// 给定一个链表: 1->2->3->4->5, 和 n = 2.

// 当删除了倒数第二个节点后，链表变为 1->2->3->5.
// 说明：

// 给定的 n 保证是有效的。

// 进阶：

// 你能尝试使用一趟扫描实现吗？

// 思路：用两个指针，第一个先跑n步，之后一起跑，第一个指针到底了，第一个指针就指向了待删除的元素，主要注意下头结点和尾节点删除的时候需要特殊处理下

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	var p2 *ListNode
	for p1 != nil {
		if n <= 0 {
			if p2 == nil {
				p2 = head
			} else {
				p2 = p2.Next
			}
		} else {
			n--
		}
		p1 = p1.Next
	}
	if p2 == nil {
		head = head.Next
	} else if p2.Next != nil {
		p2.Next = p2.Next.Next
	} else {
		head = nil
	}
	return head
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

func calcListNode(nums []int, n int) []int {
	head := createList(nums)
	head = removeNthFromEnd(head, n)
	res := printList(head)
	return res
}
